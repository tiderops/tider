package k8s

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Most PV spec fields are optional pointers that depend on the volume
// source; mapping a CSI- or NFS-backed volume used to panic on the
// HostPath and VolumeMode dereferences
func TestMapPersistentVolumeWithoutOptionalFields(t *testing.T) {
	volume := corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{Name: "pv-csi"},
		Spec: corev1.PersistentVolumeSpec{
			StorageClassName: "standard",
			AccessModes:      []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			Capacity: corev1.ResourceList{
				corev1.ResourceStorage: resource.MustParse("10Gi"),
			},
			ClaimRef: &corev1.ObjectReference{Name: "data-claim"},
			// HostPath, Local, NFS, VolumeMode all nil
		},
		Status: corev1.PersistentVolumeStatus{Phase: corev1.VolumeBound},
	}

	dto := mapPersistentVolume(volume)

	if dto.Name != "pv-csi" || dto.Status != "Bound" {
		t.Errorf("name/status = %s/%s, want pv-csi/Bound", dto.Name, dto.Status)
	}
	if dto.Capacity != "10Gi" {
		t.Errorf("capacity = %q, want 10Gi", dto.Capacity)
	}
	if dto.Claim != "data-claim" {
		t.Errorf("claim = %q, want data-claim", dto.Claim)
	}
	if dto.VolumeSpec.Host.Path != "" || dto.VolumeSpec.VolumeMode != "" {
		t.Errorf("optional fields should map to empty strings, got %+v", dto.VolumeSpec)
	}
}

func TestMapPersistentVolumeWithHostPath(t *testing.T) {
	hostPathType := corev1.HostPathDirectory
	mode := corev1.PersistentVolumeFilesystem
	volume := corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{Name: "pv-host"},
		Spec: corev1.PersistentVolumeSpec{
			VolumeMode: &mode,
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				HostPath: &corev1.HostPathVolumeSource{Path: "/data", Type: &hostPathType},
			},
		},
	}

	dto := mapPersistentVolume(volume)

	if dto.VolumeSpec.Host.Path != "/data" || dto.VolumeSpec.Host.Type != "Directory" {
		t.Errorf("host path mapping = %+v", dto.VolumeSpec.Host)
	}
	if dto.VolumeSpec.VolumeMode != "Filesystem" {
		t.Errorf("volume mode = %q, want Filesystem", dto.VolumeSpec.VolumeMode)
	}
}

// PVCs rarely set DataSource, StorageClassName can be nil, and
// VolumeAttributesClassName almost never exists; all three used to be
// dereferenced unconditionally
func TestMapPersistentVolumeClaimWithoutOptionalFields(t *testing.T) {
	claim := corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{Name: "data", Namespace: "default"},
		Spec: corev1.PersistentVolumeClaimSpec{
			VolumeName: "pv-1",
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse("5Gi"),
				},
			},
			// VolumeMode, DataSource, StorageClassName, VolumeAttributesClassName all nil
		},
		Status: corev1.PersistentVolumeClaimStatus{Phase: corev1.ClaimBound},
	}

	dto := mapPersistentVolumeClaim(claim)

	if dto.Name != "data" || dto.Status != "Bound" {
		t.Errorf("name/status = %s/%s, want data/Bound", dto.Name, dto.Status)
	}
	if dto.Size != "5Gi" {
		t.Errorf("size = %q, want 5Gi", dto.Size)
	}
	if dto.VolumeClaimSpec.DataSourceName != "" || dto.StorageClass != "" {
		t.Errorf("optional fields should map to empty strings, got %+v", dto.VolumeClaimSpec)
	}
}
