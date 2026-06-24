FROM gcr.io/oss-fuzz-base/base-builder-go@sha256:94c8ed4a88d1b48af8611ee03450a14b5245d4406ebbf1447f12c527a457035a
COPY . $SRC/atlantis
COPY .clusterfuzzlite/build.sh $SRC/build.sh
WORKDIR $SRC/atlantis
