BEGIN {
	arr[3]
	#count = 0
}

{
	if($1 == 64) {
		#if (count%200 == 0) {
		split($8, arr, "=")
		print arr[2]
		#}
		#count = count + 1
	}
}

END {

}	
