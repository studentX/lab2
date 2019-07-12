package party

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listersv1 "k8s.io/client-go/listers/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

const partySched = "partysched"

// Allowable party attires
var attires = []string{"ghoul", "goblin"}

// Scheduler represents a party scheduler.
type Scheduler struct {
	client   kubernetes.Interface
	podChan  chan *v1.Pod
	stopChan chan struct{}
	nodes    listersv1.NodeLister
}

// NewScheduler returns a new scheduler.
func NewScheduler() *Scheduler {
	s := Scheduler{
		podChan:  make(chan *v1.Pod),
		stopChan: make(chan struct{}),
	}

	return &s
}

// Run the scheduler.
func (s *Scheduler) Run(ctx context.Context) error {
	if err := s.init(); err != nil {
		return err
	}

	s.loop(ctx)

	return nil
}

func (s *Scheduler) loop(ctx context.Context) {
	log.Println("Party scheduler started....")
	for {
		select {
		case <-ctx.Done():
			close(s.stopChan)
			log.Println("Party scheduler exited....")
			return
		case p, ok := <-s.podChan:
			if !ok {
				return
			}
			s.schedule(p)
		}
	}
}

func (s *Scheduler) schedule(po *v1.Pod) {
	nodes, err := s.rank(po)
	if err != nil {
		log.Printf("Boom", err)
		return
	}
	if len(nodes) == 0 {
		log.Printf("💩  Party Pooper detected `%s. Access is denied on this cluster!", po.Name)
		s.notify(fmt.Sprintf("Party pooper detected `%s", po.Name), "FailSchedule", "Warning", po)
		return
	}

	if err := s.bind(nodes[0], po); err != nil {
		log.Println("Boom", err)
		return
	}
	s.notify("Party pod scheduled", "Scheduled", "Normal", po)
}

func (s *Scheduler) init() error {
	if err := s.initClient(); err != nil {
		return err
	}
	s.initInformers()

	return nil
}

func (s *Scheduler) initInformers() {
	factory := informers.NewSharedInformerFactory(s.client, 0)
	s.nodes = factory.Core().V1().Nodes().Lister()

	pods := factory.Core().V1().Pods()
	pods.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    s.newPodCB,
		UpdateFunc: s.podChangedCB,
	})

	factory.Start(s.stopChan)
}

func (s *Scheduler) newPodCB(o interface{}) {
	po, ok := o.(*v1.Pod)
	if !ok || !isPartyPod(po) {
		return
	}
	log.Printf("Pod added `%s", po.Name)
	s.podChan <- po
}

func (s *Scheduler) podChangedCB(_, n interface{}) {
	po, ok := n.(*v1.Pod)
	if !ok || !isPartyPod(po) {
		return
	}
	log.Printf("Pod changed `%s", po.Name)

	s.podChan <- po
}

func (s *Scheduler) initClient() error {
	var err error

	if cfg, err := restclient.InClusterConfig(); err == nil {
		s.client, err = kubernetes.NewForConfig(cfg)
		return err
	}

	path := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return err
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return err
	}

	s.client, err = kubernetes.NewForConfig(cfg)
	return err
}

func (s *Scheduler) bind(n *v1.Node, p *v1.Pod) error {
	body := map[string]interface{}{
		"target": map[string]string{
			"kind":       "Node",
			"apiVersion": "v1",
			"name":       n.ObjectMeta.Name,
		},
		"metadata": map[string]string{
			"name":      p.Name,
			"namespace": p.Namespace,
		},
	}
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bind := s.client.
		CoreV1().
		RESTClient().
		Post().
		Namespace(p.Namespace).
		Resource("bindings").
		Body(data)
	var b v1.Binding
	if err := bind.Do().Into(&b); err != nil {
		return err
	}
	log.Printf("🎊🎉  Ye! Scheduled party pod %s on node %s", p.Name, n.Name)

	return nil
}

func (s *Scheduler) notify(msg, reason, kind string, p *v1.Pod) {
	timestamp := time.Now().UTC()
	_, err := s.client.CoreV1().Events(p.Namespace).Create(&v1.Event{
		Count:          1,
		Message:        msg,
		Reason:         reason,
		LastTimestamp:  metav1.NewTime(timestamp),
		FirstTimestamp: metav1.NewTime(timestamp),
		Type:           kind,
		Source:         v1.EventSource{Component: partySched},
		InvolvedObject: v1.ObjectReference{
			Kind:      "Pod",
			Name:      p.Name,
			Namespace: p.Namespace,
			UID:       p.UID,
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: p.Name + "-",
		},
	})

	if err != nil {
		log.Println("Boom event", err)
	}
}

func (s *Scheduler) rank(p *v1.Pod) ([]*v1.Node, error) {
	var (
		candidates []*v1.Node
		attire     string
		ok         bool
	)

	if attire, ok = p.Labels["costume"]; !ok {
		return candidates, fmt.Errorf("Party pooper detected! No costume `%s on pod %s", attire, p.Name)
	}
	nn, err := s.nodes.List(labels.Everything())
	if err != nil {
		return candidates, err
	}
	for _, n := range nn {
		if n.Labels["costume"] == attire {
			candidates = append(candidates, n)
		}
	}

	return candidates, nil
}

// Helpers...

func isPartyPod(po *v1.Pod) bool {
	return po.Spec.NodeName == "" && po.Spec.SchedulerName == partySched
}

func checkAttire(costume string) bool {
	for _, c := range attires {
		if c == costume {
			return true
		}
	}

	return false
}
