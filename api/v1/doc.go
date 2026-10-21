// +groupName=example.com
// +kubebuilder:object:generate=true
// +kubebuilder:ac:generate=true
package v1

//go:generate go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.21.0 object applyconfiguration paths=.
