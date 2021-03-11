echo $1
awk -f cleanup.awk $1 > clean_data
sort -n clean_data  > sorted_data
TOTAL=$(wc -l sorted_data)
awk -v total="$TOTAL" -f cdf_gen.awk sorted_data > final_data
gnuplot gplot
