#!/bin/bash
for i in {1..10}; do curl localhost:8000/work -d name=$USER -d delay=$(expr $i % 11)s; done
