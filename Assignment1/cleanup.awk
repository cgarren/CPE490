BEGIN {
	arr[3]
}

{
	if($1 == 64) {
		split($8, arr, "=")
		#if(arr[2] <= 100) {
			print arr[2]
		#}
	}
}

END {

}	
