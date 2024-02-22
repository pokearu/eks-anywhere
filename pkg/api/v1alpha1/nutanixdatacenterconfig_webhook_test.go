package v1alpha1_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/constants"
)

func TestNutanixDatacenterConfigWebhooksValidateCreate(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	_, err := dcConf.ValidateCreate()
	g.Expect(err).To(Succeed())
}

func TestNutanixDatacenterConfigWebhooksValidateCreateReconcilePaused(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	dcConf.Annotations = map[string]string{
		"anywhere.eks.amazonaws.com/paused": "true",
	}
	_, err := dcConf.ValidateCreate()
	g.Expect(err).To(Succeed())
}

func TestNutanixDatacenterConfigWebhookValidateCreateNoCredentialRef(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	dcConf.Spec.CredentialRef = nil
	_, err := dcConf.ValidateCreate()
	g.Expect(err.Error()).To(ContainSubstring("credentialRef is required to be set to create a new NutanixDatacenterConfig"))
}

func TestNutanixDatacenterConfigWebhooksValidateCreateValidaitonFailure(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	dcConf.Spec.Endpoint = ""
	_, err := dcConf.ValidateCreate()
	g.Expect(err).To(Not(Succeed()))
}

func TestNutanixDatacenterConfigWebhooksValidateUpdate(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	_, err := dcConf.ValidateCreate()
	g.Expect(err).To(Succeed())
	newSpec := nutanixDatacenterConfig()
	newSpec.Spec.CredentialRef.Name = "new-credential"
	_, err = dcConf.ValidateUpdate(newSpec)
	g.Expect(err).To(Succeed())
}

func TestNutanixDatacenterConfigWebhooksValidateUpdateReconcilePaused(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	_, err := dcConf.ValidateCreate()
	g.Expect(err).To(Succeed())
	oldSpec := nutanixDatacenterConfig()
	oldSpec.Annotations = map[string]string{
		"anywhere.eks.amazonaws.com/paused": "true",
	}
	oldSpec.Spec.CredentialRef.Name = "new-credential"
	_, err = dcConf.ValidateUpdate(oldSpec)
	g.Expect(err).To(Succeed())
}

func TestNutanixDatacenterConfigWebhooksValidateUpdateValidationFailure(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	_, err := dcConf.ValidateCreate()
	g.Expect(err).To(Succeed())
	newSpec := nutanixDatacenterConfig()
	newSpec.Spec.Endpoint = ""
	_, err = dcConf.ValidateUpdate(newSpec)
	g.Expect(err).To(Not(Succeed()))
}

func TestNutanixDatacenterConfigWebhooksValidateUpdateInvalidOldObject(t *testing.T) {
	g := NewWithT(t)
	newConf := nutanixDatacenterConfig()
	newConf.Spec.CredentialRef = nil
	_, err := newConf.ValidateUpdate(&v1alpha1.NutanixMachineConfig{})
	g.Expect(err.Error()).To(ContainSubstring("expected a NutanixDatacenterConfig but got a *v1alpha1.NutanixMachineConfig"))
}

func TestNutanixDatacenterConfigWebhooksValidateUpdateCredentialRefRemoved(t *testing.T) {
	g := NewWithT(t)
	oldConf := nutanixDatacenterConfig()
	_, err := oldConf.ValidateCreate()
	g.Expect(err).To(Succeed())
	newConf := nutanixDatacenterConfig()
	newConf.Spec.CredentialRef = nil
	_, err = newConf.ValidateUpdate(oldConf)
	g.Expect(err.Error()).To(ContainSubstring("credentialRef cannot be removed from an existing NutanixDatacenterConfig"))
}

func TestNutanixDatacenterConfigWebhooksValidateDelete(t *testing.T) {
	g := NewWithT(t)
	dcConf := nutanixDatacenterConfig()
	_, err := dcConf.ValidateCreate()
	g.Expect(err).To(Succeed())
	_, err = dcConf.ValidateDelete()
	g.Expect(err).To(Succeed())
}

func nutanixDatacenterConfig() *v1alpha1.NutanixDatacenterConfig {
	return &v1alpha1.NutanixDatacenterConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nutanix-datacenter-config",
		},
		Spec: v1alpha1.NutanixDatacenterConfigSpec{
			Endpoint: "prism.nutanix.com",
			Port:     constants.DefaultNutanixPrismCentralPort,
			CredentialRef: &v1alpha1.Ref{
				Kind: constants.SecretKind,
				Name: constants.NutanixCredentialsName,
			},
		},
	}
}
