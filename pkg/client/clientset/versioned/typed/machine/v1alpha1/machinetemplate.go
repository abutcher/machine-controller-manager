package v1alpha1

import (
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	scheme "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineTemplatesGetter has a method to return a MachineTemplateInterface.
// A group's client should implement this interface.
type MachineTemplatesGetter interface {
	MachineTemplates(namespace string) MachineTemplateInterface
}

// MachineTemplateInterface has methods to work with MachineTemplate resources.
type MachineTemplateInterface interface {
	Create(*v1alpha1.MachineTemplate) (*v1alpha1.MachineTemplate, error)
	Update(*v1alpha1.MachineTemplate) (*v1alpha1.MachineTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MachineTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.MachineTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineTemplate, err error)
	MachineTemplateExpansion
}

// machineTemplates implements MachineTemplateInterface
type machineTemplates struct {
	client rest.Interface
	ns     string
}

// newMachineTemplates returns a MachineTemplates
func newMachineTemplates(c *MachineV1alpha1Client, namespace string) *machineTemplates {
	return &machineTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineTemplate, and returns the corresponding machineTemplate object, and an error if there is any.
func (c *machineTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineTemplate, err error) {
	result = &v1alpha1.MachineTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineTemplates that match those selectors.
func (c *machineTemplates) List(opts v1.ListOptions) (result *v1alpha1.MachineTemplateList, err error) {
	result = &v1alpha1.MachineTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineTemplates.
func (c *machineTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a machineTemplate and creates it.  Returns the server's representation of the machineTemplate, and an error, if there is any.
func (c *machineTemplates) Create(machineTemplate *v1alpha1.MachineTemplate) (result *v1alpha1.MachineTemplate, err error) {
	result = &v1alpha1.MachineTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machinetemplates").
		Body(machineTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineTemplate and updates it. Returns the server's representation of the machineTemplate, and an error, if there is any.
func (c *machineTemplates) Update(machineTemplate *v1alpha1.MachineTemplate) (result *v1alpha1.MachineTemplate, err error) {
	result = &v1alpha1.MachineTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinetemplates").
		Name(machineTemplate.Name).
		Body(machineTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineTemplate and deletes it. Returns an error if one occurs.
func (c *machineTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinetemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinetemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineTemplate.
func (c *machineTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineTemplate, err error) {
	result = &v1alpha1.MachineTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machinetemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
