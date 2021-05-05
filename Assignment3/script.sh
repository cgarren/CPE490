awk -f process.awk raw_data > processed_data
sort -n processed_data  > sorted_data
TOTAL=$(wc -l sorted_data)
echo $TOTAL
awk -v total="$TOTAL" -f cdf_gen.awk sorted_data > final_data
gnuplot gplot