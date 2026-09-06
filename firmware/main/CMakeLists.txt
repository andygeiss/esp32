# sketch.cpp is the only source. It includes the transpiled controller.ino from
# beside it, which is why nothing generated is named here. arduino-esp32 is a
# managed dependency — see idf_component.yml — and the component manager adds it
# to the requirements itself.
idf_component_register(SRCS "sketch.cpp" INCLUDE_DIRS ".")
