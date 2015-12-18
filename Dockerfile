FROM scratch
MAINTAINER jordic jordic@gmail.com
# REPO http://github.com/jordic/file_server
ADD main /
VOLUME "/tmp"
EXPOSE 8080
CMD ["/main"]
