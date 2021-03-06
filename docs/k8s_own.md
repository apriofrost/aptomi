# Your own k8s cluster 
* Make sure your k8s/Docker has enough resources available to run examples:
    * At least 4 CPU cores and 12GB RAM

# Configuration
1. Check that your k8s cluster is configured properly and has **at least one context** defined:
   ```
   kubectl config get-contexts
   ```   
   
2. Import it into Aptomi as two separate clusters *cluster-us-east* and *cluster-us-west* (corresponding to two namespaces `east` and `west` in a local k8s cluster). Make
   sure to replace `[CONTEXT_NAME]` with the name of your context:
    ```
    aptomictl login -u admin -p admin
    aptomictl gen cluster -n cluster-us-east -c [CONTEXT_NAME] -N east | aptomictl policy apply -f -
    aptomictl gen cluster -n cluster-us-west -c [CONTEXT_NAME] -N west | aptomictl policy apply -f -
    ```

Now you can move on to running the examples.
