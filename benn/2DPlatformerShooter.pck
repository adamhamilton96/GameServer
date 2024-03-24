GDPC                P                                                                         P   res://.godot/exported/133200997/export-3070c538c03ee49b7677ff960a3f5195-main.scn�      :      /�A�ܶ�
�i��&[^    T   res://.godot/exported/133200997/export-36a25e342948d0ceacc500772b5412b3-player.scn  �*      �      �{�ȃ��
'G�eUY    T   res://.godot/exported/133200997/export-401d29878c58658395a1523c1c127694-floor.scn   �      �      �0�����@�#��_�    P   res://.godot/exported/133200997/export-60d094b7efb4cf265848ac05f7208de8-wall.res�/      2      ��ln���j� ���    T   res://.godot/exported/133200997/export-61d09a3c1b01aa618cf7611fc62451ae-ladder.scn         ^      ����M�,ۛ���!��    T   res://.godot/exported/133200997/export-6350fccf6e14d2f0c891730ee1d2c6cd-bullet.scn  �            @�T���߾;r�1�!�    T   res://.godot/exported/133200997/export-7112a51661d214e7433477536e4576c8-zombie.scn  �3      �      5����њ~���&�    ,   res://.godot/global_script_class_cache.cfg  �:             ��Р�8���8~$}P�    D   res://.godot/imported/icon.svg-218a8f2b3041327d8a5756f3a245f83b.ctex�      �      �̛�*$q�*�́        res://.godot/uid_cache.bin  �>      W      69͍-�ݨI`��ŕ�       res://bullet.gd         �      E$rv9D������       res://bullet.tscn.remap �7      c       �ڦQDc?�]�F��       res://floor.tscn.remap  @8      b       �׾����A�'��xi�h       res://gun.gd�	      �      a�����E����Nwn�       res://icon.svg   ;      �      C��=U���^Qu��U3       res://icon.svg.import   `      �       ��dXAж2<
�����*       res://ladder.gd        �       �TSX��lU�+�P��       res://ladder.tscn.remap �8      c       ���r��Ǐz�1'���       res://main.tscn.remap    9      a       �J�Sw� ������       res://player.gd �$      �      /Tϐ����S�(���$       res://player.tscn.remap �9      c       ������T�?�L���       res://project.binary @      y      R?�h�p�vG��O       res://wall.tres.remap    :      a       �9���`iP����5.       res://zombie.gd �0      �      V�&�jЦM��       res://zombie.tscn.remap p:      c       V�-&�Gx��CK��d        extends PhysicsBody2D

var direction: Vector2
var source_direction: Vector2
var speed: float = 1000

var time_to_dead = 10

var dmg = 1

func _physics_process(delta) -> void:
	var collision_info = move_and_collide(((direction * speed) + source_direction) * delta)
	
	if collision_info:
		if collision_info.get_collider().name == "Zombie":
			collision_info.get_collider().hit_by_projectile(dmg)
		queue_free()
	
	time_to_dead -= delta
	
	if time_to_dead < 0:
		queue_free()
      RSRC                    PackedScene            ��������                                                  resource_local_to_scene    resource_name    custom_solver_bias    size    script 	   _bundled       Script    res://bullet.gd ��������
   Texture2D    res://icon.svg *�8CP4
      local://RectangleShape2D_1ih1l {         local://PackedScene_rk3gc �         RectangleShape2D       
     @A  �@         PackedScene          	         names "   	      CharacterBody2D    collision_layer    collision_mask    script    CollisionShape2D    shape 	   Sprite2D    scale    texture    	   variants                                       
   ��L=
ף<               node_count             nodes     !   ��������        ����                                        ����                           ����                         conn_count              conns               node_paths              editable_instances              version             RSRC     RSRC                    PackedScene            ��������                                                  resource_local_to_scene    resource_name 	   _bundled    script       Shape2D    res://wall.tres �@�q��F
   Texture2D    res://icon.svg *�8CP4
      local://PackedScene_tc61m 1         PackedScene          	         names "   	      Floor 	   position    StaticBody2D    CollisionShape2D    shape    metadata/_edit_lock_ 	   Sprite2D    scale    texture    	   variants       
     D @D                
   x�A��L>               node_count             nodes     !   ��������       ����                            ����                                 ����                               conn_count              conns               node_paths              editable_instances              version             RSRC             extends Sprite2D

var bullet_scene
var recoil_force = 500

# Called when the node enters the scene tree for the first time.
func _ready():
	bullet_scene = preload("res://bullet.tscn")

# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	
	if Input.is_action_just_pressed("Fire"):
		# spawn bullet
		var bullet_instance = bullet_scene.instantiate()
		var direction = Vector2.RIGHT.rotated(rotation)
		bullet_instance.rotation = rotation
		bullet_instance.position = global_position + direction * 20
		bullet_instance.direction = direction
		bullet_instance.source_direction = get_parent().velocity
		get_parent().get_parent().get_node("Bullets").add_child(bullet_instance)
		get_parent().shoot(-direction, recoil_force)
  GST2   �   �      ����               � �        �  RIFF�  WEBPVP8L�  /������!"2�H�$�n윦���z�x����դ�<����q����F��Z��?&,
ScI_L �;����In#Y��0�p~��Z��m[��N����R,��#"� )���d��mG�������ڶ�$�ʹ���۶�=���mϬm۶mc�9��z��T��7�m+�}�����v��ح�m�m������$$P�����එ#���=�]��SnA�VhE��*JG�
&����^x��&�+���2ε�L2�@��		��S�2A�/E���d"?���Dh�+Z�@:�Gk�FbWd�\�C�Ӷg�g�k��Vo��<c{��4�;M�,5��ٜ2�Ζ�yO�S����qZ0��s���r?I��ѷE{�4�Ζ�i� xK�U��F�Z�y�SL�)���旵�V[�-�1Z�-�1���z�Q�>�tH�0��:[RGň6�=KVv�X�6�L;�N\���J���/0u���_��U��]���ǫ)�9��������!�&�?W�VfY�2���༏��2kSi����1!��z+�F�j=�R�O�{�
ۇ�P-�������\����y;�[ ���lm�F2K�ޱ|��S��d)é�r�BTZ)e�� ��֩A�2�����X�X'�e1߬���p��-�-f�E�ˊU	^�����T�ZT�m�*a|	׫�:V���G�r+�/�T��@U�N׼�h�+	*�*sN1e�,e���nbJL<����"g=O��AL�WO!��߈Q���,ɉ'���lzJ���Q����t��9�F���A��g�B-����G�f|��x��5�'+��O��y��������F��2�����R�q�):VtI���/ʎ�UfěĲr'�g�g����5�t�ۛ�F���S�j1p�)�JD̻�ZR���Pq�r/jt�/sO�C�u����i�y�K�(Q��7őA�2���R�ͥ+lgzJ~��,eA��.���k�eQ�,l'Ɨ�2�,eaS��S�ԟe)��x��ood�d)����h��ZZ��`z�պ��;�Cr�rpi&��՜�Pf��+���:w��b�DUeZ��ڡ��iA>IN>���܋�b�O<�A���)�R�4��8+��k�Jpey��.���7ryc�!��M�a���v_��/�����'��t5`=��~	`�����p\�u����*>:|ٻ@�G�����wƝ�����K5�NZal������LH�]I'�^���+@q(�q2q+�g�}�o�����S߈:�R�݉C������?�1�.��
�ڈL�Fb%ħA ����Q���2�͍J]_�� A��Fb�����ݏ�4o��'2��F�  ڹ���W�L |����YK5�-�E�n�K�|�ɭvD=��p!V3gS��`�p|r�l	F�4�1{�V'&����|pj� ߫'ş�pdT�7`&�
�1g�����@D�˅ �x?)~83+	p �3W�w��j"�� '�J��CM�+ �Ĝ��"���4� ����nΟ	�0C���q'�&5.��z@�S1l5Z��]�~L�L"�"�VS��8w.����H�B|���K(�}
r%Vk$f�����8�ڹ���R�dϝx/@�_�k'�8���E���r��D���K�z3�^���Vw��ZEl%~�Vc���R� �Xk[�3��B��Ğ�Y��A`_��fa��D{������ @ ��dg�������Mƚ�R�`���s����>x=�����	`��s���H���/ū�R�U�g�r���/����n�;�SSup`�S��6��u���⟦;Z�AN3�|�oh�9f�Pg�����^��g�t����x��)Oq�Q�My55jF����t9����,�z�Z�����2��#�)���"�u���}'�*�>�����ǯ[����82һ�n���0�<v�ݑa}.+n��'����W:4TY�����P�ר���Cȫۿ�Ϗ��?����Ӣ�K�|y�@suyo�<�����{��x}~�����~�AN]�q�9ޝ�GG�����[�L}~�`�f%4�R!1�no���������v!�G����Qw��m���"F!9�vٿü�|j�����*��{Ew[Á��������u.+�<���awͮ�ӓ�Q �:�Vd�5*��p�ioaE��,�LjP��	a�/�˰!{g:���3`=`]�2��y`�"��N�N�p���� ��3�Z��䏔��9"�ʞ l�zP�G�ߙj��V�>���n�/��׷�G��[���\��T��Ͷh���ag?1��O��6{s{����!�1�Y�����91Qry��=����y=�ٮh;�����[�tDV5�chȃ��v�G ��T/'XX���~Q�7��+[�e��Ti@j��)��9��J�hJV�#�jk�A�1�^6���=<ԧg�B�*o�߯.��/�>W[M���I�o?V���s��|yu�xt��]�].��Yyx�w���`��C���pH��tu�w�J��#Ef�Y݆v�f5�e��8��=�٢�e��W��M9J�u�}]釧7k���:�o�����Ç����ս�r3W���7k���e�������ϛk��Ϳ�_��lu�۹�g�w��~�ߗ�/��ݩ�-�->�I�͒���A�	���ߥζ,�}�3�UbY?�Ӓ�7q�Db����>~8�]
� ^n׹�[�o���Z-�ǫ�N;U���E4=eȢ�vk��Z�Y�j���k�j1�/eȢK��J�9|�,UX65]W����lQ-�"`�C�.~8ek�{Xy���d��<��Gf�ō�E�Ӗ�T� �g��Y�*��.͊e��"�]�d������h��ڠ����c�qV�ǷN��6�z���kD�6�L;�N\���Y�����
�O�ʨ1*]a�SN�=	fH�JN�9%'�S<C:��:`�s��~��jKEU�#i����$�K�TQD���G0H�=�� �d�-Q�H�4�5��L�r?����}��B+��,Q�yO�H�jD�4d�����0*�]�	~�ӎ�.�"����%
��d$"5zxA:�U��H���H%jس{���kW��)�	8J��v�}�rK�F�@�t)FXu����G'.X�8�KH;���[             [remap]

importer="texture"
type="CompressedTexture2D"
uid="uid://kqgkey31utti"
path="res://.godot/imported/icon.svg-218a8f2b3041327d8a5756f3a245f83b.ctex"
metadata={
"vram_texture": false
}
 extends Area2D

func _on_body_entered(body):
	if body.is_in_group("Climbers"):
		if body.climbing == false:
			body.climbing = true


func _on_body_exited(body):
	if body.is_in_group("Climbers"):
		if body.climbing == true:
			body.climbing = false
       RSRC                    PackedScene            ��������                                                  resource_local_to_scene    resource_name    custom_solver_bias    size    script 	   _bundled       Script    res://ladder.gd ��������
   Texture2D    res://icon.svg *�8CP4
      local://RectangleShape2D_a43gl {         local://PackedScene_6i5fj �         RectangleShape2D       
     �A  C         PackedScene          	         names "         Area2D    script    CollisionShape2D    shape 	   Sprite2D 	   position    scale    texture    _on_body_entered    body_entered    _on_body_exited    body_exited    	   variants                           
    �����
     �>  �?               node_count             nodes        ��������        ����                            ����                           ����                               conn_count             conns                	                            
                    node_paths              editable_instances              version             RSRC  RSRC                    PackedScene            ��������                                                  ..    Player    resource_local_to_scene    resource_name 	   _bundled    script       PackedScene    res://floor.tscn y��t��Wm   PackedScene    res://player.tscn ���\TI   PackedScene    res://zombie.tscn \�z�p�)   PackedScene    res://ladder.tscn ��x��$MF      local://PackedScene_mk1w0 �         PackedScene          	         names "         Node2D    Floor    collision_layer    Floor2 	   position    scale    Floor3    Player    Bullets    Zombie    target    Ladder    z_index    	   variants                       
   \�LD  �C
   33?  �?
     dC  �C
   q=�>  �?                  
     �C  �B
   ��L>���>                         ����
    ��C �D      node_count             nodes     P   ��������        ����                ���                            ���                                        ���                                        ���                             ����                ���	                  	   
  @
               ���                               conn_count              conns               node_paths              editable_instances              version             RSRC      extends CharacterBody2D


const SPEED = 300.0
const DECELERATION = 100.0
const CLIMBING_SPEED = 100.0
const JUMP_VELOCITY = -400.0
const GUN_DIST_FROM_BODY = 20

@onready var gun = $Gun

var climbing = false
var x_force = 0
var move_force = 0
var gun_force = 0

# Get the gravity from the project settings to be synced with RigidBody nodes.
var gravity = ProjectSettings.get_setting("physics/2d/default_gravity")

func _physics_process(delta):
	# Add the gravity.
	if not is_on_floor() and not climbing:
		velocity.y += gravity * delta
	elif climbing:
		velocity.y = 0
		if Input.is_action_pressed("Up"):
			velocity.y = -CLIMBING_SPEED
		elif Input.is_action_pressed("Down"):
			velocity.y = CLIMBING_SPEED

	# Handle jump.
	if Input.is_action_just_pressed("ui_accept") and is_on_floor():
		velocity.y = JUMP_VELOCITY

	# Get the input direction and handle the movement/deceleration.
	# As good practice, you should replace UI actions with custom gameplay actions.
	var direction = Input.get_axis("Left", "Right")
	if direction:
		move_force = direction * SPEED
	else:
		move_force = move_toward(velocity.x, 0, DECELERATION)
		
	if not gun_force == 0:
		gun_force = move_toward(gun_force, 0, DECELERATION)
		
	var mouse_pos = get_local_mouse_position()
	
	gun.rotation = mouse_pos.angle()
	gun.position = mouse_pos.normalized() * GUN_DIST_FROM_BODY
	
	velocity.x = move_force + gun_force
	
	move_and_slide()
	
func shoot(direction, force):
	gun_force = direction.x * force
	velocity.y += direction.y * force
           RSRC                    PackedScene            ��������                                                  resource_local_to_scene    resource_name    custom_solver_bias    size    script 	   _bundled       Script    res://player.gd ��������
   Texture2D    res://icon.svg *�8CP4
   Script    res://gun.gd ��������      local://RectangleShape2D_qsv1u �         local://PackedScene_py276 �         RectangleShape2D       
     �A  �A         PackedScene          	         names "         CharacterBody2D    collision_mask    script 	   Climbers    CollisionShape2D    shape    metadata/_edit_lock_ 
   Character    scale    texture 	   Sprite2D    Gun 	   position 	   rotation    	   variants    
                                   
     `>" x>         
     B  ��   ���=
   ��L>���=               node_count             nodes     5   ��������        ����                                    ����                           
      ����         	                        
      ����                     	         	             conn_count              conns               node_paths              editable_instances              version             RSRC   RSRC                    RectangleShape2D            ��������                                                  resource_local_to_scene    resource_name    custom_solver_bias    size    script           local://RectangleShape2D_tllu4 �          RectangleShape2D       
     �D  �A      RSRC              extends CharacterBody2D


const SPEED = 300.0

var health = 3

# Get the gravity from the project settings to be synced with RigidBody nodes.
var gravity = ProjectSettings.get_setting("physics/2d/default_gravity")

@export var target: CharacterBody2D

func _physics_process(delta):
	# Add the gravity.
	if not is_on_floor():
		velocity.y += gravity * delta

	# Get the input direction and handle the movement/deceleration.
	# As good practice, you should replace UI actions with custom gameplay actions.
	var direction = (target.position - position).normalized().x
	if direction:
		velocity.x = direction * SPEED
	else:
		velocity.x = move_toward(velocity.x, 0, SPEED)
	
	move_and_slide()

func hit_by_projectile(dmg):
	health -= dmg
	if health <= 0:
		queue_free()
  RSRC                    PackedScene            ��������                                                  resource_local_to_scene    resource_name    custom_solver_bias    size    script 	   _bundled       Script    res://zombie.gd ��������
   Texture2D    res://icon.svg *�8CP4
      local://RectangleShape2D_0d1bv {         local://PackedScene_gpjci �         RectangleShape2D       
     �B  �B         PackedScene          	         names "         CharacterBody2D    collision_mask    script    CollisionShape2D 	   position    shape 	   Sprite2D    texture    	   variants                       
         �?                         node_count             nodes        ��������        ����                                  ����                                 ����                   conn_count              conns               node_paths              editable_instances              version             RSRC      [remap]

path="res://.godot/exported/133200997/export-6350fccf6e14d2f0c891730ee1d2c6cd-bullet.scn"
             [remap]

path="res://.godot/exported/133200997/export-401d29878c58658395a1523c1c127694-floor.scn"
              [remap]

path="res://.godot/exported/133200997/export-61d09a3c1b01aa618cf7611fc62451ae-ladder.scn"
             [remap]

path="res://.godot/exported/133200997/export-3070c538c03ee49b7677ff960a3f5195-main.scn"
               [remap]

path="res://.godot/exported/133200997/export-36a25e342948d0ceacc500772b5412b3-player.scn"
             [remap]

path="res://.godot/exported/133200997/export-60d094b7efb4cf265848ac05f7208de8-wall.res"
               [remap]

path="res://.godot/exported/133200997/export-7112a51661d214e7433477536e4576c8-zombie.scn"
             list=Array[Dictionary]([])
     <svg height="128" width="128" xmlns="http://www.w3.org/2000/svg"><rect x="2" y="2" width="124" height="124" rx="14" fill="#363d52" stroke="#212532" stroke-width="4"/><g transform="scale(.101) translate(122 122)"><g fill="#fff"><path d="M105 673v33q407 354 814 0v-33z"/><path fill="#478cbf" d="m105 673 152 14q12 1 15 14l4 67 132 10 8-61q2-11 15-15h162q13 4 15 15l8 61 132-10 4-67q3-13 15-14l152-14V427q30-39 56-81-35-59-83-108-43 20-82 47-40-37-88-64 7-51 8-102-59-28-123-42-26 43-46 89-49-7-98 0-20-46-46-89-64 14-123 42 1 51 8 102-48 27-88 64-39-27-82-47-48 49-83 108 26 42 56 81zm0 33v39c0 276 813 276 813 0v-39l-134 12-5 69q-2 10-14 13l-162 11q-12 0-16-11l-10-65H447l-10 65q-4 11-16 11l-162-11q-12-3-14-13l-5-69z"/><path d="M483 600c3 34 55 34 58 0v-86c-3-34-55-34-58 0z"/><circle cx="725" cy="526" r="90"/><circle cx="299" cy="526" r="90"/></g><g fill="#414042"><circle cx="307" cy="532" r="60"/><circle cx="717" cy="532" r="60"/></g></g></svg>
             ;x�[��a$   res://Web/index.apple-touch-icon.png�#s��	:8   res://Web/index.icon.pngo�k��Mt   res://Web/index.pngD���p1o   res://bullet.tscny��t��Wm   res://floor.tscn*�8CP4
   res://icon.svg��x��$MF   res://ladder.tscn V��F��+   res://main.tscn���\TI   res://player.tscn�@�q��F   res://wall.tres\�z�p�)   res://zombie.tscn         ECFG      application/config/name         2DPlatformShooter      application/run/main_scene         res://main.tscn    application/config/features$   "         4.2    Forward Plus       application/config/icon         res://icon.svg     input/Right0              deadzone      ?      events              InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode   D   	   key_label             unicode    d      echo          script            InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode    @ 	   key_label             unicode           echo          script      
   input/Left0              deadzone      ?      events              InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode   A   	   key_label             unicode    a      echo          script            InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode    @ 	   key_label             unicode           echo          script         input/Up0              deadzone      ?      events              InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode   W   	   key_label             unicode    w      echo          script            InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode    @ 	   key_label             unicode           echo          script      
   input/Down0              deadzone      ?      events              InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode   S   	   key_label             unicode    s      echo          script            InputEventKey         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          pressed           keycode           physical_keycode    @ 	   key_label             unicode           echo          script      
   input/Fire�              deadzone      ?      events              InputEventMouseButton         resource_local_to_scene           resource_name             device     ����	   window_id             alt_pressed           shift_pressed             ctrl_pressed          meta_pressed          button_mask           position              global_position               factor       �?   button_index         canceled          pressed           double_click          script         physics/2d/default_gravity       ��D#   rendering/renderer/rendering_method         gl_compatibility       