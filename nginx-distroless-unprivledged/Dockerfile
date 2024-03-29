ARG NGINX_VERSION
ARG DISTROLESS_BASE

FROM nginxinc/nginx-unprivileged:${NGINX_VERSION}-alpine AS base
ARG TIME_ZONE
USER root
RUN echo "************" && find /usr/share -name "zoneinfo"
RUN mkdir -p /opt/var/cache/nginx && \
    cp -a --parents /usr/lib/nginx /opt && \
    cp -a --parents /usr/share/nginx /opt && \
    cp -a --parents /var/log/nginx /opt && \
    cp -aL --parents /var/run /opt && \
    cp -a --parents /etc/nginx /opt && \
    cp -a --parents /etc/passwd /opt && \
    cp -a --parents /etc/group /opt && \
    cp -a --parents /usr/sbin/nginx /opt && \
    cp -a --parents /usr/sbin/nginx-debug /opt && \
    cp -a --parents /lib/libcrypto.so.* /opt && \
    cp -a --parents /lib/libz.* /opt && \
    cp -a --parents /lib/libc.* /opt && \
    cp -a --parents /lib/ld-musl-x86_64.so.* /opt && \
    cp -a --parents /lib/libssl.so.* /opt && \
    cp -a --parents /usr/lib/lib* /opt && \
    cp /usr/share/zoneinfo/${TIME_ZONE:-ROC} /opt/etc/localtime

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/${DISTROLESS_BASE} AS nginx-distroless-unprivileged
COPY --from=base /opt /
EXPOSE 8080
STOPSIGNAL SIGTERM
USER 1001
ENTRYPOINT ["nginx", "-g", "daemon off;"]