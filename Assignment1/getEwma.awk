BEGIN{
	RTT_NEW = 0
	RTT_OLD = 0
	count = 0
	alpha = 0.25
}

{
	if(count == 0) {
		RTT_OLD = $1
		count = 1
	}

	RTT_NEW = alpha*$1 + (1 - alpha)*RTT_OLD
	print $1, RTT_NEW
	RTT_OLD = RTT_NEW
}

END{
}
