#!/bin/basH

clusterctl alpha topology-dryrun -f $1.yaml > $1_out.yaml
