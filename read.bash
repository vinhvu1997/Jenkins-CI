#/bin/bash

INSTANCE_ID=$(cat hello.txt)
echo $INSTANCE_ID
cat hello.txt | while read line
do
        if [[ $line = *"InstanceId"* ]]
        then
		echo $line | sed 's|\"InstanceId\": \"||' | sed 's|\",||' > Instance_ID
        fi
done

