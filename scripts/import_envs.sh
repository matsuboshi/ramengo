#!/bin/bash
export $(cat .env | grep -v ^#)
