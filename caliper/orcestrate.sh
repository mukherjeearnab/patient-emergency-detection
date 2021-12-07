
# function create_config{
#     sed -e "s/RUNTPS/$1/" \
#         -e "s/SEED/$2/" \
#         ./benchmarks/benchmark_config_template.yaml
# }

echo "" > report.txt

for i in {25..300..25}
    do 
        DD = date +%s 
        echo "Running $i TPS at $DD"
        # echo "$(create_config $i $i)" > ./benchmarks/benchmark_config.yaml
        bash create_benchmark_config.sh $i $i$DD > ./benchmarks/benchmark_config.yaml
        bash ./run.sh > tempout.txt
        tail -34 tempout.txt | head -8 >> report.txt
        tail -34 tempout.txt | head -8
    done