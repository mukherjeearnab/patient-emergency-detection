sed -e "s/RUNTPS/$1/" \
    -e "s/SEED/$2/" \
    ./benchmarks/benchmark_config_template.yaml