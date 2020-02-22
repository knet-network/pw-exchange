FROM scratch

COPY ./bin/pw-exchange /bin/exchange

ENTRYPOINT ["/bin/exchange"]