#!/bin/sh

ydate=$(date -v-1d +%F)
year=${ydate:0:4}
mon=${ydate:5:2}
day=${ydate:8:2}
