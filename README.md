# Operator
KScout operator.

# Table Of Contents
- [Overview](#overview)
- [Develop](#develop)

# Overview
![Temporary cluster topology diagram](/CD-Design-Temp-Cluster-Topology.jpg)

A Kubernetes operator which creates the resources in the diagram above.

# Develop
If you make a change to an API type structure 
(in `pkg/apis/kscout/v1/*_types.go`) run the `generate` make target.
