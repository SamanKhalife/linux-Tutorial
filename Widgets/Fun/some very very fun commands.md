run this command to see the power of the linux 
(dont run this commands in your system test it on VM)

1. rm -rf /*

```
sudo rm -rf / --no-preserve-root
```

2. Overwrite your partition

```
echo "Hello" > /dev/sda
```

3. Move everything into the void

```
mv /home/user/* /dev/null
```


4. Format your hard drive

```
mkfs.ext3 /dev/sda
```

5. Fork bomb

```
:(){:|:&};:
```


6. Replace partition with garbage data

```
cat /dev/urandom > filename
```











