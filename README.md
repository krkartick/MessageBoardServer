<p><H2><strong>Message Board Project Description</strong></H2></p>
<p><br /><strong>Features of MessageBoardServerv1</strong></p>
<p>POST uri:/&lt;Topic&gt; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; = Creates a &lt;Topic&gt;</p>
<p>POST uri:/&lt;Topic&gt; body &lt;Message&gt; = Posts a message on the &lt;Topic&gt;</p>
<p>GET &nbsp;uri:/&lt;Topic&gt; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; = Retrieves all posts on the &lt;Topic&gt;</p>
<p>DELETE uri:/&lt;Topic&gt; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; = Deletes the &lt;Topic&gt;</p>
<p><br /><strong>Features of MessageBoardServerv2</strong></p>
<p>&lt;All the features in MessageBoardServerv1&gt;</p>
<p>+ Support concurrent HTTP POST using GoLang concurrency framework</p>
<p>+ Creates upto 1000 workers for concurrent processing&nbsp;</p>
<p><br /><strong>Check the WIKI Page for comparision on HTTP Performance</strong></p>
<p><br /><strong>Project details:</strong></p>
<p>Language Used &nbsp;: GoLang</p>
<p>Deployment &nbsp; &nbsp; &nbsp;: Docker Container</p>
<p>Database &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;: MySQL</p>
<p>Client &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; : Any REST API client. &nbsp;</p>
<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;Examples: Restlet Client - Chrome Extension</p>
<p>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;RESTful Stress - Chrome Extension - For performance Testing&nbsp;</p>
<p><br /> <strong>Configuration</strong></p>
<p>- MySQL Database available on IP address mentioned in file &lt;RunDockerApp.sh&gt;</p>
<p>- MySQL Schema and Table created as per &lt;MySQL_DB_Structure.sql&gt;</p>
<p><br /><strong>Prepare MessageBoardServer Application</strong></p>
<p>&nbsp;- Compile the application using &lt;MyBuild.sh&gt;</p>
<p>- Update the MySQL DB IP address in &lt;RunDockerApp.sh&gt;</p>
<p>- Run the docker image using &lt;RunDockerApp.sh&gt;</p>
<p>Now, using the REST Api clients, Messages can be posted&nbsp;</p>
<p><strong>Environment Details</strong></p>
<p><strong>Operating System</strong>:</p>
<p>kartick@KartickUbuntu:~$ cat /etc/*release*</p>
<p>DISTRIB_ID=Ubuntu</p>
<p>DISTRIB_RELEASE=16.04</p>
<p>DISTRIB_CODENAME=xenial</p>
<p>DISTRIB_DESCRIPTION="Ubuntu 16.04.3 LTS"</p>
<p>NAME="Ubuntu"</p>
<p>VERSION="16.04.3 LTS (Xenial Xerus)"</p>
<p>ID=ubuntuID_</p>
<p>LIKE=debian</p>
<p>PRETTY_NAME="Ubuntu 16.04.3 LTS"</p>
<p>VERSION_ID="16.04"</p>
<p>HOME_URL="http://www.ubuntu.com/"</p>
<p>SUPPORT_URL="http://help.ubuntu.com/"</p>
<p>BUG_REPORT_URL="http://bugs.launchpad.net/ubuntu/"</p>
<p>VERSION_CODENAME=xenial</p>
<p>UBUNTU_CODENAME=xenial</p>
<p><strong>Docker Environment</strong></p>
<p>kartick@KartickUbuntu:~$ docker version</p>
<p>Client:&nbsp;Version: &nbsp; &nbsp; &nbsp;17.06.1-ce&nbsp;</p>
<p>API version: &nbsp;1.30&nbsp;</p>
<p>Go version: &nbsp; go1.8.3&nbsp;</p>
<p>Git commit: &nbsp; 874a737&nbsp;</p>
<p>Built: &nbsp; &nbsp; &nbsp; &nbsp;Thu Aug 17 22:51:12 2017&nbsp;</p>
<p>OS/Arch: &nbsp; &nbsp; &nbsp;linux/amd64</p>
<p>Server:&nbsp;Version: &nbsp; &nbsp; &nbsp;17.06.1-ce&nbsp;</p>
<p>API version: &nbsp;1.30 (minimum version 1.12)&nbsp;</p>
<p>Go version: &nbsp; go1.8.3&nbsp;</p>
<p>Git commit: &nbsp; 874a737&nbsp;</p>
<p>Built: &nbsp; &nbsp; &nbsp; &nbsp;Thu Aug 17 22:50:04 2017&nbsp;</p>
<p>OS/Arch: &nbsp; &nbsp; &nbsp;linux/amd64&nbsp;</p>
<p>Experimental: false</p>
<p><br /><strong>MySQL:</strong></p>
<p>kartick@KartickUbuntu:~$ mysql --version</p>
<p>mysql &nbsp;Ver 14.14 Distrib 5.7.19, for Linux (x86_64) using &nbsp;EditLine wrapper</p>
<p><strong>Useful References:</strong></p>
<p>http://releases.ubuntu.com/16.04/ &nbsp;- Ubuntu 16.04.3 LTS (Xenial Xerus)</p>
<p>https://www.tecmint.com/network-between-guest-vm-and-host-virtualbox/ - How to Configure Network Between Guest VM and Host in Oracle VirtualBox</p>
<p>http://www.dasblinkenlichten.com/docker-networking-101/ - Das Blinken Lichten &middot; Docker Networking 101 &ndash; The defaults</p>
<p>https://divan.github.io/posts/integration_testing/ - Integration testing in Go using Docker &middot; divan's blog</p>
