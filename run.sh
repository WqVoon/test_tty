#!/bin/bash

stty cbreak -echo && go run . && stty -cbreak echo