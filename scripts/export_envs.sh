#!/bin/sh
export $(cat .env | grep -v ^#)
