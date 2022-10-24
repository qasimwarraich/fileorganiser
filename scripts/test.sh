#!/bin/bash


mkdir test_dir

touch -d 20200201 test_dir/2020-02-01.txt
touch -d 20210420 test_dir/2021-04-20.txt
touch -d 20220201 test_dir/2022-02-01.txt

mkdir test_dir/2022-02-01
touch -d 20220201 test_dir/2022-02-01/

mkdir test_dir/2022-02-02
touch -d 20220202 test_dir/2022-02-02/

mkdir test_dir/randdir
touch -d 20220202 test_dir/randdir/
