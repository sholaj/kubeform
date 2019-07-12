package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleContainerAnalysisNote struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleContainerAnalysisNoteSpec   `json:"spec,omitempty"`
	Status            GoogleContainerAnalysisNoteStatus `json:"status,omitempty"`
}

type GoogleContainerAnalysisNoteSpecAttestationAuthorityHint struct {
	HumanReadableName string `json:"human_readable_name"`
}

type GoogleContainerAnalysisNoteSpecAttestationAuthority struct {
	Hint []GoogleContainerAnalysisNoteSpecAttestationAuthority `json:"hint"`
}

type GoogleContainerAnalysisNoteSpec struct {
	Name                 string                            `json:"name"`
	Project              string                            `json:"project"`
	AttestationAuthority []GoogleContainerAnalysisNoteSpec `json:"attestation_authority"`
}

type GoogleContainerAnalysisNoteStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleContainerAnalysisNoteList is a list of GoogleContainerAnalysisNotes
type GoogleContainerAnalysisNoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleContainerAnalysisNote CRD objects
	Items []GoogleContainerAnalysisNote `json:"items,omitempty"`
}