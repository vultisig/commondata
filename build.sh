#!/bin/sh
buf format -w && buf lint && buf build && buf generate