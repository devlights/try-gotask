#!/usr/bin/env bash

# 
# 指定していない場合は .task となる
#
task

#
# TASK_TEMP_DIR 環境変数を指定。.task2 ディレクトリにチェックサムを生成するように指定。
# (通常は export で設定してしまうのが多い)
#
TASK_TEMP_DIR='.task2' task

#
# --temp-dirオプションで指定 (v3.52.0で追加)
#
task --temp-dir='.task3'
