<template>
    <div ref="sceneContainer"></div>
  </template>
  
  <script>
  import * as THREE from 'three';
  import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';
  import { OrbitControls } from 'three/addons/controls/OrbitControls.js';
  
  export default {
    name: 'ThreeScene',
    mounted() {
      this.initRenderer();
      this.initScene();
      this.initCamera();
      this.initControls();
      this.loadModel();
      this.animate();
      window.addEventListener('resize', this.onWindowResize);
    },
    beforeUnmount() {
      window.removeEventListener('resize', this.onWindowResize);
    },
    methods: {
      initRenderer() {
        this.renderer = new THREE.WebGLRenderer({ antialias: true });
        this.renderer.outputColorSpace = THREE.SRGBColorSpace;
        this.renderer.setSize(window.innerWidth, window.innerHeight);
        this.renderer.setClearColor(0x000000);
        this.renderer.setPixelRatio(window.devicePixelRatio);
        this.renderer.shadowMap.enabled = true;
        this.renderer.shadowMap.type = THREE.PCFSoftShadowMap;
        this.$refs.sceneContainer.appendChild(this.renderer.domElement);
      },
      initScene() {
        this.scene = new THREE.Scene();
  
        const groundGeometry = new THREE.PlaneGeometry(20, 20, 32, 32);
        groundGeometry.rotateX(-Math.PI / 2);
        const groundMaterial = new THREE.MeshStandardMaterial({
          color: 0x555555,
          side: THREE.DoubleSide
        });
        const groundMesh = new THREE.Mesh(groundGeometry, groundMaterial);
        groundMesh.castShadow = false;
        groundMesh.receiveShadow = true;
        this.scene.add(groundMesh);
  
        const spotLight = new THREE.SpotLight(0xffffff, 3000, 100, 0.22, 1);
        spotLight.position.set(0, 25, 0);
        spotLight.castShadow = true;
        spotLight.shadow.bias = -0.0001;
        this.scene.add(spotLight);
      },
      initCamera() {
        this.camera = new THREE.PerspectiveCamera(45, window.innerWidth / window.innerHeight, 1, 1000);
        this.camera.position.set(4, 5, 11);
      },
      initControls() {
        this.controls = new OrbitControls(this.camera, this.renderer.domElement);
        this.controls.enableDamping = true;
        this.controls.enablePan = false;
        this.controls.minDistance = 5;
        this.controls.maxDistance = 20;
        this.controls.minPolarAngle = 0.5;
        this.controls.maxPolarAngle = 1.5;
        this.controls.autoRotate = false;
        this.controls.target = new THREE.Vector3(0, 1, 0);
        this.controls.update();
      },
      loadModel() {
        const loader = new GLTFLoader().setPath('public/low-poly-m18-hellcat/');
        loader.load('model.gltf', (gltf) => {
          console.log('loading model');
          const mesh = gltf.scene;
  
          mesh.traverse((child) => {
            if (child.isMesh) {
              child.castShadow = true;
              child.receiveShadow = true;
            }
          });
  
          mesh.position.set(0, 1.05, -1);
          this.scene.add(mesh);
  
          document.getElementById('progress-container').style.display = 'none';
        }, (xhr) => {
          console.log(`loading ${xhr.loaded / xhr.total * 100}%`);
        }, (error) => {
          console.error(error);
        });
      },
      animate() {
        requestAnimationFrame(this.animate);
        this.controls.update();
        this.renderer.render(this.scene, this.camera);
      },
      onWindowResize() {
        this.camera.aspect = window.innerWidth / window.innerHeight;
        this.camera.updateProjectionMatrix();
        this.renderer.setSize(window.innerWidth, window.innerHeight);
      }
    }
  }
  </script>