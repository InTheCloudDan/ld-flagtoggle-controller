# LaunchDarkly Kubernetes FlagToggle Controller

This is a reference implementation for a kubernetes controller to toggle LaunchDarkly flags via the API.

##Requirements

Go 1.12+ installed
[Kubebuilder v2.0.0-alpha.4](https://github.com/kubernetes-sigs/kubebuilder/releases/tag/v2.0.0-alpha.4) - The binaries that are part of the release must be under `/usr/local/kubebuilder/bin`.

## Setup
You need a local or remote kubernetes cluster running. For local you can use [kind](https://github.com/kubernetes-sigs/kind)

1. Add the CRD definition to your cluster: `kubectly apply -f config/crd/bases/featureflags.launchdarkly.com_flagtoggles.yaml`
2. You need to have a LaunchDarkly API Key with `write` permissions. Once you get the key you can add it as an environment variable in your current environment with the key `LD_API_KEY`
2. `make run` will start the controller on your local machine and connect to kubernetes based on your `KUBECONFIG` environment variable. 


`sample-crd.yaml` is a sample that you can modify to match a field in your environment.
