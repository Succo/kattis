while diff smart.ans slow.ans > /dev/null; 
do
	generate > sample-03.in
	cat sample-03.in | 10_types_people > smart.ans
	cat sample-03.in | slow_10 > slow.ans
	diff smart.ans slow.ans
done
