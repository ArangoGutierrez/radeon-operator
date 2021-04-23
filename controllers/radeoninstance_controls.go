/*
Copyright Eduardo Arango 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	secv1 "github.com/openshift/api/security/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type controlFunc []func(n RI) (ResourceStatus, error)

type ResourceStatus int

const (
	Ready ResourceStatus = iota
	NotReady
)

func (s ResourceStatus) String() string {
	names := [...]string{
		"Ready",
		"NotReady"}

	if s < Ready || s > NotReady {
		return "Unkown Resources Status"
	}
	return names[s]
}

func Namespace(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].Namespace

	found := &corev1.Namespace{}
	logger := log.WithValues("Namespace", obj.Name, "Namespace", "Cluster")

	logger.Info("Looking for")
	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating ")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, skipping update")

	return Ready, nil
}

func ServiceAccount(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].ServiceAccount

	obj.SetNamespace(r.ins.GetNamespace())

	found := &corev1.ServiceAccount{}
	logger := log.WithValues("ServiceAccount", obj.Name, "Namespace", obj.Namespace)

	logger.Info("Looking for")

	if err := controllerutil.SetControllerReference(r.ins, &obj, r.rec.Scheme); err != nil {
		return NotReady, err
	}

	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating ")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, skipping update")

	return Ready, nil
}

func ClusterRole(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].ClusterRole

	found := &rbacv1.ClusterRole{}
	logger := log.WithValues("ClusterRole", obj.Name, "Namespace", obj.Namespace)

	logger.Info("Looking for")

	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: "", Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = r.rec.Client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func ClusterRoleBinding(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].ClusterRoleBinding

	found := &rbacv1.ClusterRoleBinding{}
	logger := log.WithValues("ClusterRoleBinding", obj.Name, "Namespace", obj.Namespace)

	obj.Subjects[0].Namespace = r.ins.GetNamespace()

	logger.Info("Looking for")

	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: "", Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = r.rec.Client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func Role(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].Role

	obj.SetNamespace(r.ins.GetNamespace())

	found := &rbacv1.Role{}
	logger := log.WithValues("Role", obj.Name, "Namespace", obj.Namespace)

	logger.Info("Looking for")

	if err := controllerutil.SetControllerReference(r.ins, &obj, r.rec.Scheme); err != nil {
		return NotReady, err
	}

	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = r.rec.Client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func RoleBinding(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].RoleBinding

	obj.SetNamespace(r.ins.GetNamespace())

	found := &rbacv1.RoleBinding{}
	logger := log.WithValues("RoleBinding", obj.Name, "Namespace", obj.Namespace)

	logger.Info("Looking for")

	if err := controllerutil.SetControllerReference(r.ins, &obj, r.rec.Scheme); err != nil {
		return NotReady, err
	}

	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = r.rec.Client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func DaemonSet(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].DaemonSet

	// update the image
	obj.Spec.Template.Spec.Containers[0].Image = r.ins.Spec.ImagePath()

	// update image pull policy
	if r.ins.Spec.ImagePullPolicy != "" {
		obj.Spec.Template.Spec.Containers[0].ImagePullPolicy = r.ins.Spec.ImagePolicy(r.ins.Spec.ImagePullPolicy)
	}

	obj.SetNamespace(r.ins.Namespace)

	found := &appsv1.DaemonSet{}
	logger := log.WithValues("DaemonSet", obj.Name, "Namespace", obj.Namespace)

	logger.Info("Looking for")

	if err := controllerutil.SetControllerReference(r.ins, &obj, r.rec.Scheme); err != nil {
		return NotReady, err
	}

	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = r.rec.Client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func SecurityContextConstraints(r RI) (ResourceStatus, error) {

	state := r.idx
	obj := r.resources[state].SecurityContextConstraints

	// Set the correct namespace for SCC when installed in non default namespace
	obj.Users[0] = "system:serviceaccount:" + r.ins.GetNamespace() + ":" + obj.GetName()

	found := &secv1.SecurityContextConstraints{}
	logger := log.WithValues("SecurityContextConstraints", obj.Name, "Namespace", "default")

	logger.Info("Looking for")

	err := r.rec.Client.Get(context.TODO(), types.NamespacedName{Namespace: "", Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = r.rec.Client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create", "Error", err)
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")

	required := obj.DeepCopy()
	required.ResourceVersion = found.ResourceVersion

	err = r.rec.Client.Update(context.TODO(), required)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}
