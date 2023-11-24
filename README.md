# pfe-shamir-poc

Small proof of concept for my final school project.

it contain 2 types of golang webservices :
- Manager : manage storage services and dispatch secrets using shamir secret sharing algorithm
- Storage : store and serve secrets when manager request for it

It is modular as you can configure how many parts/servers you want to split your secret, and how many parts you need to retrieve the whole secret.

dont hesitate to ask for a demo !
