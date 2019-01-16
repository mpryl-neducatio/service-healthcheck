#!/bin/bash

docker run -it --rm --name "service-healthcheck" -v $(pwd):/project service-healthcheck:latest bash
