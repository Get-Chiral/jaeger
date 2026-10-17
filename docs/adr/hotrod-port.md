# HotROD example uses port 8090 by default

Status: proposed (Chiral /chiral confirm test, do not merge)

The HotROD example listens on port 8090 by default, because port 8080 collides with the Jaeger query service when both run on one machine.
