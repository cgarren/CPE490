echo $1
awk -f cleanup.awk $1 > clean_data
#sort -n clean_data  > sorted_data
#TOTAL=$(wc -l sorted_data)
#awk -v total="$TOTAL" -f cdf_gen.awk sorted_data > final_data
awk -f makePretty.awk clean_data > final_data
awk -f getEwma.awk clean_data > rtt_data
gnuplot gplot
