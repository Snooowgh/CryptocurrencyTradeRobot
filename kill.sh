NAME="go"  
echo "Start trying to kill $NAME process"  
ID=`ps -ef | grep "$NAME" | grep -v "$0" | grep -v "grep" | awk '{print $2}'`  
echo $ID  
for id in $ID  
do  
kill -9 $id  
echo "kill $id"  
done  
echo  "Done"
