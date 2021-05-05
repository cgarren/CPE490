BEGIN{
	count = 0
	MAX = total
}

{
	count++
	cumul_cdf = count/MAX
	print $1, cumul_cdf
}

END{
}
