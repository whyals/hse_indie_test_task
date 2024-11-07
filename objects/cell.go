components {
  id: "circleLogic"
  component: "/scripts/cell_func.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"gray_circle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/assets.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.072
    y: 0.072
    z: 1.0E-6
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 35.0\n"
  "}\n"
  ""
}
