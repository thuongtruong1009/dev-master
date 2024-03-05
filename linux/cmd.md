### Hierarchy

![](./linux_hierarchy.png)

1. **/bin**: Essential user command binaries

2. **/boot**: Static files of the boot loader

3. **/dev**: Device files

4. **/etc**: Host-specific system-wide configuration files

5. **/home**: Users' home directories

6. **/lib**: Essential shared libraries and kernel modules

7. **/media**: Mount point for removable media

8. **/mnt**: Mount point for mounting a filesystem temporarily

9. **/opt**: Add-on application software packages

10. **/proc**: Virtual filesystem providing process and kernel information

11. **/root**: Home directory for the root user

12. **/run**: Run-time variable data

13. **/sbin**: Essential system binaries

14. **/srv**: Data for services provided by this system

15. **/sys**: Virtual filesystem providing kernel information

16. **/tmp**: Temporary files

17. **/usr**: Secondary hierarchy

18. **/var**: Variable data

### Commands

1. **clear**: Clear terminal

2. **pwd**: Current path

3. **whoami**: Current user

4. **ls**: List files

    - **-a**: display hidden files
    - **-l**: display as list
    - **-h**: Human readable
    - **-t**: Sort by time (newest first)
    - **-r**: Reverse sort
    - **-R**: Recursive
    - **-S**: Sort by size
    - **-1**: One file per line

5. **cd**: Change directory

6. **touch**: Create file

7. **mkdir**: Create directory

    - **-p**: Create parent directories if they don't exist

8. **rm**: Remove file

    - **-r**: Remove directory
    - **-f**: Force remove

9. **rmdir**: Remove directory

10. **cp**: Copy file

11. **mv**: Move file

ex: `mv file1/a.txt file2/b.txt` will move file1 to file2 and rename a.txt to b.txt

12. **cat**: View content files

13. **more**: Display file

14. **less**: Display file

15. **head**: Display first lines of file

16. **tail**: Display last lines of file

-   **-n**: Number of lines

ex: `tail -n 5 file.txt > log.txt` will display last 5 lines of file.txt and write to log.txt

-   **-f**: Follow file content tail real-time

17. **grep**: Search file

18. **find**: Find file

19. **locate**: Locate file

20. **which**: Locate a command

21. **whereis**: Locate a command

22. **echo**: Print to terminal

ex: `echo "Hello World"`

ex: `echo "Hello World" > file.txt` will write "Hello World" to file.txt

ex: `echo "Hello World" >> file.txt` will append "Hello World" to file.txt

23. **history**: view history of commands

24. **sudo**: Run command as superuser

ex: `sudo apt-get update`

-   **-i**: enter superuser mode

25. **shutdown**: Shutdown system

26. **chown**: Change owner

27. **su**: Switch user

28. **chmod**: Change permissions

29. **passwd**: Change password

30. **useradd**: Add user

31. **userdel**: Delete user

32. **groupadd**: Add group

33. **groupdel**: Delete group

34. **usermod**: Modify user

35. **groupmod**: Modify group

36. **chgrp**: Change group

37. **ps**: Process status

    -   **-e**: All processes
    -   **-f**: Full format
    -   **-l**: Long format
    -   **-u**: User processes

38. **kill**: Kill process

39. **top**: Display system information

40. **df**: Disk free

41. **du**: Disk usage

42. **free**: Memory usage

43. **uname**: System information

44. **date**: Display date

45. **cal**: Display calendar

46. **time**: Measure time

47. **tar**: Archive files

    -   **-c**: Create
    -   **-x**: Extract
    -   **-v**: Verbose
    -   **-f**: File
    -   **-z**: Compress
    -   **-j**: Bzip2
    -   **-t**: List
    -   **-r**: Append
    -   **-u**: Update
    -   **-A**: Concatenate
    -   **-d**: Diff
    -   **-W**: Verify

ex: `tar -cvf file.tar file1 file2` will create file.tar with file1 and file2

ex: `tar -xvf file.tar` will extract file.tar

48. **gzip**: Compress file

    -   **-d**: Decompress

49. **bzip2**: Compress file

    -   **-d**: Decompress

50. **zip**: Compress file

    -   **-r**: Recursive

51. **unzip**: Decompress file

52. **wget**: Download file

53. **curl**: Transfer data

54. **ssh**: Secure shell

55. **scp**: Secure copy

56. **rsync**: Remote sync

57. **ping**: Test network connection

ex: `ping google.com`

58. **traceroute**: Trace network route

-   **-T**: TCP
-   **-U**: UDP
-   **-p**: Port

ex: `traceroute -T 192.168.1.9`

59. **netstat**: Network statistics

60. **ifconfig**: Network interface configuration

61. **route**: Network routing table

62. **hostname**: Hostname

63. **nslookup**: Name server lookup

64. **dig**: Domain information groper

65. **host**: DNS lookup

66. **whois**: Who is

67. **lsof**: List open files

68. **ps**: Process status

69. **kill**: Kill process

70. **nice**: Set priority

71. **renice**: Change priority

72. **jobs**: List jobs

73. **bg**: Background job

74. **fg**: Foreground job

75. **nohup**: Run command immune to hangups

76. **at**: Schedule command

77. **cron**: Schedule command

78. **watch**: Execute command periodically

79. **alias**: Create alias

80. **unalias**: Remove alias

81. **source**: Execute file

82. **exit**: Exit terminal

83. **logout**: Logout

84. **reboot**: Reboot system

85. **apt install**: Install package

86. **apt remove**: Remove package

### File Types

1. **-**: Regular file

2. **d**: Directory

3. **l**: Symbolic link

4. **c**: Character device

5. **b**: Block device

6. **p**: Named pipe

7. **s**: Socket

### Permissions

1. **r**: Read

2. **w**: Write

3. **x**: Execute

4. **-**: No permission

### Users

1. **root**: Superuser

2. **user**: Regular user

### Groups

1. **root**: Superuser group

2. **user**: Regular user group
