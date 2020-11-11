FROM erichough/nfs-server:2.2.1

COPY entrypoint.sh /usr/local/bin/

COPY fake-nfsserver /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

CMD ["/usr/local/bin/fake-nfsserver"]
