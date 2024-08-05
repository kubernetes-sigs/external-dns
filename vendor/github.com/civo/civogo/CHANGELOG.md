
0.2.57
=============
2021-11-02

* Remove implicit call to templates (#63) (ff81332)

0.2.54
=============
2021-09-27

* Provide region when reboot, stop and start instance (#62) (e04b7763)

0.2.52
=============
2021-08-30

* Add new struct fields for cluster firewall (#60) (d69977c6)
* Update the changelog (6f7d2d14)

0.2.51
=============
2021-08-20

* Add firewall ID to instance config (for creating instance) (#59) (63eb8ece)

0.2.50
=============
2021-08-20

* Exclude disk images with name containing "k3s" (#58) (4a2e2643)
* Fix lint issues (51b3d43b)
* Fixing capitalisation of decodeERROR to go standards and UnknowError typo (cd1195df)
* Add endpoints for team/permission management (e0bec13d)

0.2.49
=============
2021-07-16

* Remove unnecessary code from the error handler file and add new error handler (14fc34c7)

0.2.48
=============
2021-07-15

* Change de default error in case the error is not in the list (1bf187fd)
* Added a new error handler (cf996f92)

0.2.47
=============
2021-05-23

* Fixed error in the recycle kuberntes node (d9a041df)
* Updated the Change log (267f348f)

0.2.46
=============
2021-05-21

* Added omitempty to some models, to prevent show empty result (f99bbb6a)
* Fixed test for the instance (63305185)

0.2.45
=============
2021-05-10

* Fixed error in the instance controller (a9f4749f)
* Added new error handler (bfb7b690)

0.2.44
=============
2021-05-05

* Merge branch 'master' of https://github.com/civo/civogo (7f556d7b)
* Added the pool option to kubernestes object (2459ccd5)

0.2.43
=============
2021-05-04

* Added change to kubernetes and volume (0f3b7fb6)
* Add cluster_id to volume creation (47dce610)
* Fix CHANGELOG (7a917e8e)
* Add GetVolume call to save always returning then filtering down to a single volume (419cda25)
* Add Status to volume (4884e404)
* Added region to volume attach/detach (093df491)
* Update CHANGELOG for 0.2.38 (9c1df9d1)
* Update VolumeConfig to have a namespace attribute (89995b6a)
* Update fake_client.go (d629abae)

0.2.37
=============
2021-04-08

* Added new methods for DiskImage (cd0f7879)

0.2.36
=============
2021-04-06

* Added omitempty to kubernetes config, that way we only send the filled fields (af03bb1f)
* Updated the changelog (dc3ea1df)

0.2.35
=============
2021-04-02

* Merge branch 'master' of https://github.com/civo/civogo (772d0a16)
* Fixed error in the code, added a new handler for network error (c47419ce)
* Updated the changelog (ed132dec)
* Added new fields to kubernetes struct (43e5e9a8)
* Updated the changelog (699e5392)
* Merge branch 'master' of https://github.com/civo/civogo (1e9c4c70)
* Fixed error in the struct for InstalledApplication in kubernetes (ff8a577c)
* Updated the changelog (0cf2bd6e)
* Added network to the firewall struct (f173564c)
* Added GetDefaultRegion to the region (e2bd46f7)
* Updated the changelog (9c6173fb)
* Added default field to the region (79f15aa1)
* Updated the changelog (b45dffe1)
* Feature/add region (#50) (19e61587)
* Clusternames should be considered case insensitive (#49) (dfcbfaec)
* Update Change log (c237384c)
* Fixed check in handler for errors (c5ae6fb3)
* Update Change log (5564e235)
* Fixed lint check in the actions (346fe0ff)
* Fixed some bugs in the error handler (a0333a83)
* Update Change log (ce2a8a20)
* Added option to use proxy if is present in the system (3c3da053)
* Update Change log (d6c63e63)
* Fixed the recycle option in Kubernetes (405567d0)
* Update Change log (4b69a797)
* Updated all find function to all command (8ac5d69e)
* Update Change log (dd52c163)
* Fixed an error in the struct of the creation of intances (e7cd9f3b)
* Update Change log (ede2e4c1)
* Added two more error to the error handler (c1e74133)
* Fixed error in the DNS test (b0996760)
* Update Change log (d6191cc6)
* Added new record type to DNS (6e3d7b1e)
* Update Change log (fc9cd8ff)
* improved error handling in the library (7a0e384f)
* Update test.yml (0aa9f6bf)
* Update test.yml (73500add)
* Update test.yml (f3d121d8)
* Update test.yml (2df11af6)
* Update Change log (c024bec2)
* Updated all test (d8b73535)
* Update Change log (63a3ae56)
* Added constant errors to the lib (4095fdfc)
* Update the chnage log (b11ac276)
* Added FindTemplate fucntion to the template module (59a86819)
* Update the change log (4cc1b488)
* Revert "Update the change log" (21b44c0e)
* Update the change log (8cfac7bc)
* Add constant errors (#41) (c983d56e)
* Added CPU, RAM and SSD fields to Instance struct (f6135e29)
* Added new feature (b74d3224)
* Fixed error in the cluster upgrade cmd (c49389ab)
* Add the new UpgradeAvailableTo field to KubernetesCluster (9c392181)
* Change application struct in the kubernetes module (#39) (59b86eba)
* Change application struct in kubernetes (c1839b96)
* added new way to search in network (923b509b)
* feat(kubernetes): new added option at the moment scaling down the cluster (#35) (1906a5fc)
* Add pagination for Kubernetes clusters (#34) (6ce671a8)
* (hotfix) change snapshot config (77d29967)
* Change PublicIPRequired to a string to support IP moving (d0635c7e)
* add template endpoints (c73d51fa)
* Minor tweaks to SSH key struct (d20f49e0)
* update the ssh key file (f5eab5e2)
* Add new option to ssh (fdf845a6)
* Fail the build if any lint issues (95ae34b1)
* add rename endpoint for firewall (588340b6)
* Updated the lib to use json and not form for all request (09617776)
* Add body to client error responses (a8ebeb95)
* Merge pull request #25 from alejandrojnm/hostfix/add_function_firewall (5ccea39d)
* Merge pull request #24 from alejandrojnm/hotfix/urgent_fix (7b07ee24)
* Change GetDNSRecord to find by ID not name (4f53b021)
* Remove DNSDomainID from DNSRecordConfig (edfa5583)
* Merge pull request #23 from alejandrojnm/feature/update_dns_record (d8c0e967)
* Fix instance test (5ae22ab1)
* Fix signature of SetInstanceTags (c470a7d5)
* Add GetInstanceConsoleURL method (1032a3be)
* Check links in go documentation (efbe8dbc)
* Change link to docs (83fdce86)
* Update badges (7cbcb38c)
* Fix tests after some minor refactoring (f149ccdc)
* Change GB to Gigabytes (f0fa3a15)
* Fix lint issues (62a82aec)
* Change workflow to custom action (143ac314)
* Move lint to its own action (541c21a6)
* Use user-contributed lint action (4e265ae4)
* Install golint before running it (0aa650f4)
* Add lint to tests and Go 1.14 (0bbdcd0c)
* Add changelog (f2cd42c5)
* Add Find* methods to all areas (d9686526)
* Fix some linting issues with Kubernetes (8dafd66a)
* Add LICENSE (dd906190)
* Return an error if creating a client with an empty API key (ee8dab3b)
* Remove progress table from README as its now module complete (1bfa0af0)
* Add minor tweaks to Webhooks (0471e7d2)
* Merge pull request #21 from alejandrojnm/feature/webhook (2d7d45db)
* Change minor naming issues in snapshot module (fc4a75b7)
* Merge pull request #20 from alejandrojnm/feature/api_snapshots (4b2ae45c)
* Fix a couple of comments (81e4652e)
* Add charge endpoint (457dfa5d)
* Change awkward names of units in field names (f20cab75)
* Add instance size endpoints (971a4a47)
* Add quota endpoint (7b54da2c)
* Add region endpoint (66b590cc)
* Change List functions to be plural and prefixing DNS (e3433bb0)
* Update display of README progress (43d6bc08)
* Merge pull request #19 from alejandrojnm/hotfix/loadbalance_readme (eacb35ca)
* Merge pull request #18 from alejandrojnm/master (41c7acb3)
* Merge pull request #16 from alejandrojnm/add-kubernetes-apis (730f2dd4)
* Add minor changes to Load Balancer methods (b8d5ddb8)
* Merge pull request #14 from alejandrojnm/feature/loadbalancer (e69c9746)
* Merge pull request #11 from rberrelleza/get-record (5590adda)
* Merge pull request #12 from rberrelleza/add-owners (338c4dfd)
* Add client.LastJSONResponse for the CLI to use (5378d41d)
* Merge pull request #7 from rberrelleza/fix-casing-forms (b3aba767)
* Merge pull request #9 from alejandrojnm/hostfix-network (f66cbcbb)
* Change string to map of strings (ab4ebaa8)
* Cleanup some minor issues (a69bef1d)
* Merge pull request #3 from rberrelleza/add-dns-apis (34870810)
* Merge pull request #4 from alejandrojnm/add-firewall-apis (554cc2b3)
* Merge pull request #5 from alejandrojnm/add-volumes-apis (9e8048fc)
* Merge pull request #6 from alejandrojnm/add-network-options (bbb04c3d)
* Change order of automated tests (57fae691)
* Ensure gofmt code compliance (e6d1cdcb)
* Remove panic (5d9bf4a7)
* Remove string interpolation, more go idiomatic this way (50fd2f10)
* Move JSON parsing for SimpleResponse to a helper method (8e8caf0a)
* Add CONTRIBUTING guide (a6d63538)
* Add status bage to README (1eabf12e)
* Add Go module support (2e043e15)
* Add GitHub action to test golang (e9c40745)
* Fix tests (08e7668c)
* Add Instance functions (264d4c87)
* Add Instance methods,first version (cd217b8b)
* Add tests for instance (31634d0f)
* Add Instance listing/fetching/creating (dc16cca8)
* - Add more options to the network options (5e258256)
* - Add New network options (6fbeb8a8)
* Add option to create a new network (e4a70af4)
* Fix error in volumes.go in the VolumesConfig (3db7ca46)
* Add volumes option to the civogo (52dca353)
* - Fix a error in firewall.go (5aecfa1c)
* - Fix some error in firewall.go (070605de)
* Add firewall option to the civogo (572557b6)
* fix test (ebde1c12)
* handle http errors (9c0f0967)
* add tests (37733dc1)
* Add functions to manipulate records (5f324bc9)
* - Fix a error network, now you can rename a network (79674176)
* Merge pull request #1 from civo/master (cbeb0cd4)
* change config for `form` (1ea969af)
* Add a default owner for the repo (fb92d21f)
* use Create instead of New (3055381e)
* New and delete Domain (5c6fe06b)
* add a result const (9c28c61c)
* fix GetDomain test (9df135c1)
* get and list records (5bf08246)
* update and new domain (acb6cd6a)
* - Update name in loadbalancer_test.go (4048f9c1)
* - Update the loadbalancer_test.go (fea3cf5b)
* - Fix some error in types in loadbalancer.go and loadbalancer_test.go (84b247e5)
* - Now can add, delete update and list the load balance (e2e8858b)
* - Initial files (4002274e)
* - Fix all recommendation from Andy, change NumTargetNodes from string to int, and in UpdateKubernetesClusters change map[string]string in favor of map[string]interface{} (2fdaa159)
* - Fix all recommendation from Andy (fa672090)
* - Final change in the kubernetes interface (7e0221f0)
* - Some change in kubernetes (9bb817f9)
* - First commit for kubernetes (05711724)
* - Update README.md with a list of all modules (58429347)
* Merge pull request #3 from civo/master (5a3cfcd6)
* Merge pull request #2 from civo/master (311f778e)
* - Update README.md, Fix error in loadbalancer.go (74877fa8)
* - Done the snapshot module (d2ede4b0)
* - First part of the snapshot api (eaf506c7)
* Merge branch 'master' into feature/api_snapshots (2408aaa3)
* Merge pull request #4 from civo/master (6d43cd65)
* Merge pull request #5 from civo/master (6e2137e8)
* - Add webhook module (467bd9a5)
* Merge pull request #6 from civo/master (55ee0710)
* - Fix some bug in dns.go dns_test.go (2edd5469)
* - Add option to update dns record (8ff8a494)
* Merge pull request #7 from civo/master (08c1c198)
* fix(loadbalancer): Fix data in test (29eab040)
* fix(loadbalancer): Fix struct in loadbalancer (450e0188)
* fix(firewall): Fix error in firewall test (f8658cf3)
* fix(firewall): Fix struct in firewall (6b4ca011)
* - Fix some bug in firewall.go and firewall_test.go (7f13e1c8)
* Merge pull request #9 from civo/master (27235987)
* Merge pull request #8 from civo/master (d118d875)
* fix(firewall): Change the id passed to the function (aff66676)
* fix(firewall): Fix name in the firewall (c480c528)
* fix(firewall): Fix name in the firewall (874f9ef6)
* feat(firewall): Add new option to search rule (58a86032)
* Merge pull request #11 from civo/master (53dc6a1e)
* Merge pull request #10 from alejandrojnm/hotfix/urgent_fix (99662c54)
* updated README.md (#51) (fce968e2)

0.2.34
=============
2021-03-16

* Fix more cases were we were updating a by-value array (a2afa167)

0.2.33
=============
2021-03-16

* Finding a volume by ID should work in the fake client (1f6ff05c)

0.2.32
=============
2021-03-11

* Update changelog (863f3d0d)
* Add simple fake tests to ensure it conforms to the interface and the pattern works (06e89dda)

0.2.31
=============
2021-03-11

* Fix error in NewFirewall definition for FakeClient (a5721041)

0.2.30
=============
2021-03-10

* Add fake client (#53) (318bdea6)
* Updated the changelog (ed132dec)
* Added new fields to kubernetes struct (43e5e9a8)
* Updated the changelog (699e5392)
* Merge branch 'master' of https://github.com/civo/civogo (1e9c4c70)
* Fixed error in the struct for InstalledApplication in kubernetes (ff8a577c)
* Updated the changelog (0cf2bd6e)
* Added network to the firewall struct (f173564c)
* Added GetDefaultRegion to the region (e2bd46f7)
* Updated the changelog (9c6173fb)
* Added default field to the region (79f15aa1)
* Updated the changelog (b45dffe1)
* Feature/add region (#50) (19e61587)
* Clusternames should be considered case insensitive (#49) (dfcbfaec)
* Update Change log (c237384c)
* Fixed check in handler for errors (c5ae6fb3)
* Update Change log (5564e235)
* Fixed lint check in the actions (346fe0ff)
* Fixed some bugs in the error handler (a0333a83)
* Update Change log (ce2a8a20)
* Added option to use proxy if is present in the system (3c3da053)
* Update Change log (d6c63e63)
* Fixed the recycle option in Kubernetes (405567d0)
* Update Change log (4b69a797)
* Updated all find function to all command (8ac5d69e)
* Update Change log (dd52c163)
* Fixed an error in the struct of the creation of intances (e7cd9f3b)
* Update Change log (ede2e4c1)
* Added two more error to the error handler (c1e74133)
* Fixed error in the DNS test (b0996760)
* Update Change log (d6191cc6)
* Added new record type to DNS (6e3d7b1e)
* Update Change log (fc9cd8ff)
* improved error handling in the library (7a0e384f)
* Update test.yml (0aa9f6bf)
* Update test.yml (73500add)
* Update test.yml (f3d121d8)
* Update test.yml (2df11af6)
* Update Change log (c024bec2)
* Updated all test (d8b73535)
* Update Change log (63a3ae56)
* Added constant errors to the lib (4095fdfc)
* Update the chnage log (b11ac276)
* Added FindTemplate fucntion to the template module (59a86819)
* Update the change log (4cc1b488)
* Revert "Update the change log" (21b44c0e)
* Update the change log (8cfac7bc)
* Add constant errors (#41) (c983d56e)
* Added CPU, RAM and SSD fields to Instance struct (f6135e29)
* Added new feature (b74d3224)
* Fixed error in the cluster upgrade cmd (c49389ab)
* Add the new UpgradeAvailableTo field to KubernetesCluster (9c392181)
* Change application struct in the kubernetes module (#39) (59b86eba)
* Change application struct in kubernetes (c1839b96)
* added new way to search in network (923b509b)
* feat(kubernetes): new added option at the moment scaling down the cluster (#35) (1906a5fc)
* Add pagination for Kubernetes clusters (#34) (6ce671a8)

0.2.29
=============
2021-02-28

* Updated the changelog (ed132dec)
* Added new fields to kubernetes struct (43e5e9a8)

0.2.28
=============
2021-02-26

* Updated the changelog (699e5392)
* Merge branch 'master' of https://github.com/civo/civogo (1e9c4c70)
* Fixed error in the struct for InstalledApplication in kubernetes (ff8a577c)
* Updated the changelog (0cf2bd6e)

0.2.27
=============
2021-02-10

* Added network to the firewall struct (f173564c)

0.2.26
=============
2021-02-09

* Added GetDefaultRegion to the region (e2bd46f7)
* Updated the changelog (9c6173fb)

0.2.25
=============
2021-02-09

* Added default field to the region (79f15aa1)
* Updated the changelog (b45dffe1)

0.2.24
=============
2021-02-08

* Feature/add region (#50) (19e61587)
* Clusternames should be considered case insensitive (#49) (dfcbfaec)

0.2.23
=============
2020-12-04

* Update Change log (c237384c)
* Fixed check in handler for errors (c5ae6fb3)

0.2.22
=============
2020-11-18

* Update Change log (5564e235)
* Fixed lint check in the actions (346fe0ff)
* Fixed some bugs in the error handler (a0333a83)

0.2.21
=============
2020-10-31

* Update Change log (ce2a8a20)
* Added option to use proxy if is present in the system (3c3da053)

0.2.20
=============
2020-10-31

* Update Change log (d6c63e63)
* Fixed the recycle option in Kubernetes (405567d0)
* Update Change log (4b69a797)
* Updated all find function to all command (8ac5d69e)
* Update Change log (dd52c163)
* Fixed an error in the struct of the creation of intances (e7cd9f3b)
* Update Change log (ede2e4c1)
* Added two more error to the error handler (c1e74133)
* Fixed error in the DNS test (b0996760)
* Update Change log (d6191cc6)
* Added new record type to DNS (6e3d7b1e)
* Update Change log (fc9cd8ff)
* improved error handling in the library (7a0e384f)
* Update test.yml (0aa9f6bf)
* Update test.yml (73500add)
* Update test.yml (f3d121d8)
* Update test.yml (2df11af6)
* Update Change log (c024bec2)
* Updated all test (d8b73535)
* Update Change log (63a3ae56)
* Added constant errors to the lib (4095fdfc)
* Update the chnage log (b11ac276)
* Added FindTemplate fucntion to the template module (59a86819)
* Update the change log (4cc1b488)
* Revert "Update the change log" (21b44c0e)
* Update the change log (8cfac7bc)
* Add constant errors (#41) (c983d56e)
* Added CPU, RAM and SSD fields to Instance struct (f6135e29)
* Added new feature (b74d3224)
* Fixed error in the cluster upgrade cmd (c49389ab)
* Add the new UpgradeAvailableTo field to KubernetesCluster (9c392181)
* Change application struct in the kubernetes module (#39) (59b86eba)
* Change application struct in kubernetes (c1839b96)
* added new way to search in network (923b509b)
* feat(kubernetes): new added option at the moment scaling down the cluster (#35) (1906a5fc)
* Add pagination for Kubernetes clusters (#34) (6ce671a8)
* (hotfix) change snapshot config (77d29967)

0.2.19
=============
2020-09-25

* Update Change log (dd52c163)
* Fixed an error in the struct of the creation of intances (e7cd9f3b)

0.2.18
=============
2020-09-23

* Update Change log (ede2e4c1)
* Added two more error to the error handler (c1e74133)
* Fixed error in the DNS test (b0996760)

0.2.17
=============
2020-09-13

* Update Change log (d6191cc6)
* Added new record type to DNS (6e3d7b1e)

0.2.16
=============
2020-08-24

* Update Change log (fc9cd8ff)
* improved error handling in the library (7a0e384f)

0.2.15
=============
2020-08-17

* Update test.yml (0aa9f6bf)
* Update test.yml (73500add)
* Update test.yml (f3d121d8)
* Update test.yml (2df11af6)
* Update Change log (c024bec2)
* Updated all test (d8b73535)
* Update Change log (63a3ae56)
* Added constant errors to the lib (4095fdfc)

0.2.14
=============
2020-08-11

* Update the chnage log (b11ac276)
* Added FindTemplate fucntion to the template module (59a86819)
* Update the change log (4cc1b488)
* Revert "Update the change log" (21b44c0e)
* Update the change log (8cfac7bc)

0.2.13
=============
2020-07-31

* Add constant errors (#41) (c983d56e)

0.2.12
=============
2020-07-07

* Added CPU, RAM and SSD fields to Instance struct (f6135e29)

0.2.11
=============
2020-07-06

* Added new feature (b74d3224)

0.2.10
=============
2020-06-24

* Fixed error in the cluster upgrade cmd (c49389ab)
* Add the new UpgradeAvailableTo field to KubernetesCluster (9c392181)
* Change application struct in the kubernetes module (#39) (59b86eba)
* Change application struct in kubernetes (c1839b96)
* added new way to search in network (923b509b)
* feat(kubernetes): new added option at the moment scaling down the cluster (#35) (1906a5fc)
* Add pagination for Kubernetes clusters (#34) (6ce671a8)
* (hotfix) change snapshot config (77d29967)
* Change PublicIPRequired to a string to support IP moving (d0635c7e)

0.2.9
=============
2020-06-19

* Add the new UpgradeAvailableTo field to KubernetesCluster (9c392181)

0.2.8
=============
2020-05-12

* Change application struct in the kubernetes module (#39) (59b86eba)

0.2.7
=============
2020-05-11

* Change application struct in kubernetes (c1839b96)

0.2.6
=============
2020-05-06

* added new way to search in network (923b509b)

0.2.5
=============
2020-04-15

* feat(kubernetes): new added option at the moment scaling down the cluster (#35) (1906a5fc)

0.2.4
=============
2020-04-07

* Add pagination for Kubernetes clusters (#34) (6ce671a8)
* (hotfix) change snapshot config (77d29967)
* Change PublicIPRequired to a string to support IP moving (d0635c7e)
* add template endpoints (c73d51fa)
* Minor tweaks to SSH key struct (d20f49e0)
* update the ssh key file (f5eab5e2)
* Add new option to ssh (fdf845a6)
* Fail the build if any lint issues (95ae34b1)
* add rename endpoint for firewall (588340b6)
* Updated the lib to use json and not form for all request (09617776)
* Add body to client error responses (a8ebeb95)
* Merge pull request #25 from alejandrojnm/hostfix/add_function_firewall (5ccea39d)
* Merge pull request #24 from alejandrojnm/hotfix/urgent_fix (7b07ee24)
* Change GetDNSRecord to find by ID not name (4f53b021)
* Remove DNSDomainID from DNSRecordConfig (edfa5583)
* Merge pull request #23 from alejandrojnm/feature/update_dns_record (d8c0e967)
* Fix instance test (5ae22ab1)
* Fix signature of SetInstanceTags (c470a7d5)
* Add GetInstanceConsoleURL method (1032a3be)
* Check links in go documentation (efbe8dbc)
* Change link to docs (83fdce86)
* Update badges (7cbcb38c)
* Fix tests after some minor refactoring (f149ccdc)
* Change GB to Gigabytes (f0fa3a15)
* Fix lint issues (62a82aec)
* Change workflow to custom action (143ac314)
* Move lint to its own action (541c21a6)
* Use user-contributed lint action (4e265ae4)
* Install golint before running it (0aa650f4)
* Add lint to tests and Go 1.14 (0bbdcd0c)
* Add changelog (f2cd42c5)
* Add Find* methods to all areas (d9686526)
* Fix some linting issues with Kubernetes (8dafd66a)
* Add LICENSE (dd906190)
* Return an error if creating a client with an empty API key (ee8dab3b)
* Remove progress table from README as its now module complete (1bfa0af0)
* Add minor tweaks to Webhooks (0471e7d2)
* Merge pull request #21 from alejandrojnm/feature/webhook (2d7d45db)
* Change minor naming issues in snapshot module (fc4a75b7)
* Merge pull request #20 from alejandrojnm/feature/api_snapshots (4b2ae45c)
* Fix a couple of comments (81e4652e)
* Add charge endpoint (457dfa5d)
* Change awkward names of units in field names (f20cab75)
* Add instance size endpoints (971a4a47)
* Add quota endpoint (7b54da2c)
* Add region endpoint (66b590cc)
* Change List functions to be plural and prefixing DNS (e3433bb0)
* Update display of README progress (43d6bc08)
* Merge pull request #19 from alejandrojnm/hotfix/loadbalance_readme (eacb35ca)
* Merge pull request #18 from alejandrojnm/master (41c7acb3)
* Merge pull request #16 from alejandrojnm/add-kubernetes-apis (730f2dd4)
* Add minor changes to Load Balancer methods (b8d5ddb8)
* Merge pull request #14 from alejandrojnm/feature/loadbalancer (e69c9746)
* Merge pull request #11 from rberrelleza/get-record (5590adda)
* Merge pull request #12 from rberrelleza/add-owners (338c4dfd)
* Add client.LastJSONResponse for the CLI to use (5378d41d)
* Merge pull request #7 from rberrelleza/fix-casing-forms (b3aba767)
* Merge pull request #9 from alejandrojnm/hostfix-network (f66cbcbb)
* Change string to map of strings (ab4ebaa8)
* Cleanup some minor issues (a69bef1d)
* Merge pull request #3 from rberrelleza/add-dns-apis (34870810)
* Merge pull request #4 from alejandrojnm/add-firewall-apis (554cc2b3)
* Merge pull request #5 from alejandrojnm/add-volumes-apis (9e8048fc)
* Merge pull request #6 from alejandrojnm/add-network-options (bbb04c3d)
* Change order of automated tests (57fae691)
* Ensure gofmt code compliance (e6d1cdcb)
* Remove panic (5d9bf4a7)
* Remove string interpolation, more go idiomatic this way (50fd2f10)
* Move JSON parsing for SimpleResponse to a helper method (8e8caf0a)
* Add CONTRIBUTING guide (a6d63538)
* Add status bage to README (1eabf12e)
* Add Go module support (2e043e15)
* Add GitHub action to test golang (e9c40745)
* Fix tests (08e7668c)
* Add Instance functions (264d4c87)
* Add Instance methods,first version (cd217b8b)
* Add tests for instance (31634d0f)
* Add Instance listing/fetching/creating (dc16cca8)
* - Add more options to the network options (5e258256)
* - Add New network options (6fbeb8a8)
* Add option to create a new network (e4a70af4)
* Fix error in volumes.go in the VolumesConfig (3db7ca46)
* Add volumes option to the civogo (52dca353)
* - Fix a error in firewall.go (5aecfa1c)
* - Fix some error in firewall.go (070605de)
* Add firewall option to the civogo (572557b6)
* fix test (ebde1c12)
* handle http errors (9c0f0967)
* add tests (37733dc1)
* Add functions to manipulate records (5f324bc9)
* - Fix a error network, now you can rename a network (79674176)
* Merge pull request #1 from civo/master (cbeb0cd4)
* change config for `form` (1ea969af)
* Add a default owner for the repo (fb92d21f)
* use Create instead of New (3055381e)
* New and delete Domain (5c6fe06b)
* add a result const (9c28c61c)
* fix GetDomain test (9df135c1)
* get and list records (5bf08246)
* update and new domain (acb6cd6a)
* - Update name in loadbalancer_test.go (4048f9c1)
* - Update the loadbalancer_test.go (fea3cf5b)
* - Fix some error in types in loadbalancer.go and loadbalancer_test.go (84b247e5)
* - Now can add, delete update and list the load balance (e2e8858b)
* - Initial files (4002274e)
* - Fix all recommendation from Andy, change NumTargetNodes from string to int, and in UpdateKubernetesClusters change map[string]string in favor of map[string]interface{} (2fdaa159)
* - Fix all recommendation from Andy (fa672090)
* - Final change in the kubernetes interface (7e0221f0)
* - Some change in kubernetes (9bb817f9)
* - First commit for kubernetes (05711724)
* - Update README.md with a list of all modules (58429347)
* Merge pull request #3 from civo/master (5a3cfcd6)
* Merge pull request #2 from civo/master (311f778e)
* - Update README.md, Fix error in loadbalancer.go (74877fa8)
* - Done the snapshot module (d2ede4b0)
* - First part of the snapshot api (eaf506c7)
* Merge branch 'master' into feature/api_snapshots (2408aaa3)
* Merge pull request #4 from civo/master (6d43cd65)
* Merge pull request #5 from civo/master (6e2137e8)
* - Add webhook module (467bd9a5)
* Merge pull request #6 from civo/master (55ee0710)
* - Fix some bug in dns.go dns_test.go (2edd5469)
* - Add option to update dns record (8ff8a494)
* Merge pull request #7 from civo/master (08c1c198)
* fix(loadbalancer): Fix data in test (29eab040)
* fix(loadbalancer): Fix struct in loadbalancer (450e0188)
* fix(firewall): Fix error in firewall test (f8658cf3)
* fix(firewall): Fix struct in firewall (6b4ca011)
* - Fix some bug in firewall.go and firewall_test.go (7f13e1c8)
* Merge pull request #9 from civo/master (27235987)
* Merge pull request #8 from civo/master (d118d875)
* fix(firewall): Change the id passed to the function (aff66676)
* fix(firewall): Fix name in the firewall (c480c528)
* fix(firewall): Fix name in the firewall (874f9ef6)
* feat(firewall): Add new option to search rule (58a86032)
* Merge pull request #11 from civo/master (53dc6a1e)
* Merge pull request #10 from alejandrojnm/hotfix/urgent_fix (99662c54)

0.2.3
=============
2020-04-03

* (hotfix) change snapshot config (77d29967)
* Change PublicIPRequired to a string to support IP moving (d0635c7e)
* add template endpoints (c73d51fa)
* Minor tweaks to SSH key struct (d20f49e0)
* update the ssh key file (f5eab5e2)
* Add new option to ssh (fdf845a6)
* Fail the build if any lint issues (95ae34b1)
* add rename endpoint for firewall (588340b6)
* Updated the lib to use json and not form for all request (09617776)
* Add body to client error responses (a8ebeb95)
* Merge pull request #25 from alejandrojnm/hostfix/add_function_firewall (5ccea39d)
* Merge pull request #24 from alejandrojnm/hotfix/urgent_fix (7b07ee24)
* Change GetDNSRecord to find by ID not name (4f53b021)
* Remove DNSDomainID from DNSRecordConfig (edfa5583)
* Merge pull request #23 from alejandrojnm/feature/update_dns_record (d8c0e967)
* Fix instance test (5ae22ab1)
* Fix signature of SetInstanceTags (c470a7d5)
* Add GetInstanceConsoleURL method (1032a3be)
* Check links in go documentation (efbe8dbc)
* Change link to docs (83fdce86)
* Update badges (7cbcb38c)
* Fix tests after some minor refactoring (f149ccdc)
* Change GB to Gigabytes (f0fa3a15)
* Fix lint issues (62a82aec)
* Change workflow to custom action (143ac314)
* Move lint to its own action (541c21a6)
* Use user-contributed lint action (4e265ae4)
* Install golint before running it (0aa650f4)
* Add lint to tests and Go 1.14 (0bbdcd0c)
* Add changelog (f2cd42c5)
* Add Find* methods to all areas (d9686526)
* Fix some linting issues with Kubernetes (8dafd66a)
* Add LICENSE (dd906190)
* Return an error if creating a client with an empty API key (ee8dab3b)
* Remove progress table from README as its now module complete (1bfa0af0)
* Add minor tweaks to Webhooks (0471e7d2)
* Merge pull request #21 from alejandrojnm/feature/webhook (2d7d45db)
* Change minor naming issues in snapshot module (fc4a75b7)
* Merge pull request #20 from alejandrojnm/feature/api_snapshots (4b2ae45c)
* Fix a couple of comments (81e4652e)
* Add charge endpoint (457dfa5d)
* Change awkward names of units in field names (f20cab75)
* Add instance size endpoints (971a4a47)
* Add quota endpoint (7b54da2c)
* Add region endpoint (66b590cc)
* Change List functions to be plural and prefixing DNS (e3433bb0)
* Update display of README progress (43d6bc08)
* Merge pull request #19 from alejandrojnm/hotfix/loadbalance_readme (eacb35ca)
* Merge pull request #18 from alejandrojnm/master (41c7acb3)
* Merge pull request #16 from alejandrojnm/add-kubernetes-apis (730f2dd4)
* Add minor changes to Load Balancer methods (b8d5ddb8)
* Merge pull request #14 from alejandrojnm/feature/loadbalancer (e69c9746)
* Merge pull request #11 from rberrelleza/get-record (5590adda)
* Merge pull request #12 from rberrelleza/add-owners (338c4dfd)
* Add client.LastJSONResponse for the CLI to use (5378d41d)
* Merge pull request #7 from rberrelleza/fix-casing-forms (b3aba767)
* Merge pull request #9 from alejandrojnm/hostfix-network (f66cbcbb)
* Change string to map of strings (ab4ebaa8)
* Cleanup some minor issues (a69bef1d)
* Merge pull request #3 from rberrelleza/add-dns-apis (34870810)
* Merge pull request #4 from alejandrojnm/add-firewall-apis (554cc2b3)
* Merge pull request #5 from alejandrojnm/add-volumes-apis (9e8048fc)
* Merge pull request #6 from alejandrojnm/add-network-options (bbb04c3d)
* Change order of automated tests (57fae691)
* Ensure gofmt code compliance (e6d1cdcb)
* Remove panic (5d9bf4a7)
* Remove string interpolation, more go idiomatic this way (50fd2f10)
* Move JSON parsing for SimpleResponse to a helper method (8e8caf0a)
* Add CONTRIBUTING guide (a6d63538)
* Add status bage to README (1eabf12e)
* Add Go module support (2e043e15)
* Add GitHub action to test golang (e9c40745)
* Fix tests (08e7668c)
* Add Instance functions (264d4c87)
* Add Instance methods,first version (cd217b8b)
* Add tests for instance (31634d0f)
* Add Instance listing/fetching/creating (dc16cca8)
* - Add more options to the network options (5e258256)
* - Add New network options (6fbeb8a8)
* Add option to create a new network (e4a70af4)
* Fix error in volumes.go in the VolumesConfig (3db7ca46)
* Add volumes option to the civogo (52dca353)
* - Fix a error in firewall.go (5aecfa1c)
* - Fix some error in firewall.go (070605de)
* Add firewall option to the civogo (572557b6)
* fix test (ebde1c12)
* handle http errors (9c0f0967)
* add tests (37733dc1)
* Add functions to manipulate records (5f324bc9)
* - Fix a error network, now you can rename a network (79674176)
* Merge pull request #1 from civo/master (cbeb0cd4)
* change config for `form` (1ea969af)
* Add a default owner for the repo (fb92d21f)
* use Create instead of New (3055381e)
* New and delete Domain (5c6fe06b)
* add a result const (9c28c61c)
* fix GetDomain test (9df135c1)
* get and list records (5bf08246)
* update and new domain (acb6cd6a)
* - Update name in loadbalancer_test.go (4048f9c1)
* - Update the loadbalancer_test.go (fea3cf5b)
* - Fix some error in types in loadbalancer.go and loadbalancer_test.go (84b247e5)
* - Now can add, delete update and list the load balance (e2e8858b)
* - Initial files (4002274e)
* - Fix all recommendation from Andy, change NumTargetNodes from string to int, and in UpdateKubernetesClusters change map[string]string in favor of map[string]interface{} (2fdaa159)
* - Fix all recommendation from Andy (fa672090)
* - Final change in the kubernetes interface (7e0221f0)
* - Some change in kubernetes (9bb817f9)
* - First commit for kubernetes (05711724)
* - Update README.md with a list of all modules (58429347)
* Merge pull request #3 from civo/master (5a3cfcd6)
* Merge pull request #2 from civo/master (311f778e)
* - Update README.md, Fix error in loadbalancer.go (74877fa8)
* - Done the snapshot module (d2ede4b0)
* - First part of the snapshot api (eaf506c7)
* Merge branch 'master' into feature/api_snapshots (2408aaa3)
* Merge pull request #4 from civo/master (6d43cd65)
* Merge pull request #5 from civo/master (6e2137e8)
* - Add webhook module (467bd9a5)
* Merge pull request #6 from civo/master (55ee0710)
* - Fix some bug in dns.go dns_test.go (2edd5469)
* - Add option to update dns record (8ff8a494)
* Merge pull request #7 from civo/master (08c1c198)
* fix(loadbalancer): Fix data in test (29eab040)
* fix(loadbalancer): Fix struct in loadbalancer (450e0188)
* fix(firewall): Fix error in firewall test (f8658cf3)
* fix(firewall): Fix struct in firewall (6b4ca011)
* - Fix some bug in firewall.go and firewall_test.go (7f13e1c8)
* Merge pull request #9 from civo/master (27235987)
* Merge pull request #8 from civo/master (d118d875)
* fix(firewall): Change the id passed to the function (aff66676)
* fix(firewall): Fix name in the firewall (c480c528)
* fix(firewall): Fix name in the firewall (874f9ef6)
* feat(firewall): Add new option to search rule (58a86032)
* Merge pull request #11 from civo/master (53dc6a1e)
* Merge pull request #10 from alejandrojnm/hotfix/urgent_fix (99662c54)

0.2.2
=============
2020-03-27

* Change PublicIPRequired to a string to support IP moving (d0635c7e)
* add template endpoints (c73d51fa)
* Minor tweaks to SSH key struct (d20f49e0)
* update the ssh key file (f5eab5e2)
* Add new option to ssh (fdf845a6)
* Fail the build if any lint issues (95ae34b1)
* add rename endpoint for firewall (588340b6)
* Updated the lib to use json and not form for all request (09617776)
* Add body to client error responses (a8ebeb95)
* Merge pull request #25 from alejandrojnm/hostfix/add_function_firewall (5ccea39d)
* Merge pull request #24 from alejandrojnm/hotfix/urgent_fix (7b07ee24)
* Change GetDNSRecord to find by ID not name (4f53b021)
* Remove DNSDomainID from DNSRecordConfig (edfa5583)
* Merge pull request #23 from alejandrojnm/feature/update_dns_record (d8c0e967)
* Fix instance test (5ae22ab1)
* Fix signature of SetInstanceTags (c470a7d5)
* Add GetInstanceConsoleURL method (1032a3be)
* Check links in go documentation (efbe8dbc)
* Change link to docs (83fdce86)
* Update badges (7cbcb38c)
* Fix tests after some minor refactoring (f149ccdc)
* Change GB to Gigabytes (f0fa3a15)
* Fix lint issues (62a82aec)
* Change workflow to custom action (143ac314)
* Move lint to its own action (541c21a6)
* Use user-contributed lint action (4e265ae4)
* Install golint before running it (0aa650f4)
* Add lint to tests and Go 1.14 (0bbdcd0c)
* Add changelog (f2cd42c5)
* Add Find* methods to all areas (d9686526)
* Fix some linting issues with Kubernetes (8dafd66a)
* Add LICENSE (dd906190)
* Return an error if creating a client with an empty API key (ee8dab3b)
* Remove progress table from README as its now module complete (1bfa0af0)
* Add minor tweaks to Webhooks (0471e7d2)
* Merge pull request #21 from alejandrojnm/feature/webhook (2d7d45db)
* Change minor naming issues in snapshot module (fc4a75b7)
* Merge pull request #20 from alejandrojnm/feature/api_snapshots (4b2ae45c)
* Fix a couple of comments (81e4652e)
* Add charge endpoint (457dfa5d)
* Change awkward names of units in field names (f20cab75)
* Add instance size endpoints (971a4a47)
* Add quota endpoint (7b54da2c)
* Add region endpoint (66b590cc)
* Change List functions to be plural and prefixing DNS (e3433bb0)
* Update display of README progress (43d6bc08)
* Merge pull request #19 from alejandrojnm/hotfix/loadbalance_readme (eacb35ca)
* Merge pull request #18 from alejandrojnm/master (41c7acb3)
* Merge pull request #16 from alejandrojnm/add-kubernetes-apis (730f2dd4)
* Add minor changes to Load Balancer methods (b8d5ddb8)
* Merge pull request #14 from alejandrojnm/feature/loadbalancer (e69c9746)
* Merge pull request #11 from rberrelleza/get-record (5590adda)
* Merge pull request #12 from rberrelleza/add-owners (338c4dfd)
* Add client.LastJSONResponse for the CLI to use (5378d41d)
* Merge pull request #7 from rberrelleza/fix-casing-forms (b3aba767)
* Merge pull request #9 from alejandrojnm/hostfix-network (f66cbcbb)
* Change string to map of strings (ab4ebaa8)
* Cleanup some minor issues (a69bef1d)
* Merge pull request #3 from rberrelleza/add-dns-apis (34870810)
* Merge pull request #4 from alejandrojnm/add-firewall-apis (554cc2b3)
* Merge pull request #5 from alejandrojnm/add-volumes-apis (9e8048fc)
* Merge pull request #6 from alejandrojnm/add-network-options (bbb04c3d)
* Change order of automated tests (57fae691)
* Ensure gofmt code compliance (e6d1cdcb)
* Remove panic (5d9bf4a7)
* Remove string interpolation, more go idiomatic this way (50fd2f10)
* Move JSON parsing for SimpleResponse to a helper method (8e8caf0a)
* Add CONTRIBUTING guide (a6d63538)
* Add status bage to README (1eabf12e)
* Add Go module support (2e043e15)
* Add GitHub action to test golang (e9c40745)
* Fix tests (08e7668c)
* Add Instance functions (264d4c87)
* Add Instance methods,first version (cd217b8b)
* Add tests for instance (31634d0f)
* Add Instance listing/fetching/creating (dc16cca8)
* - Add more options to the network options (5e258256)
* - Add New network options (6fbeb8a8)
* Add option to create a new network (e4a70af4)
* Fix error in volumes.go in the VolumesConfig (3db7ca46)
* Add volumes option to the civogo (52dca353)
* - Fix a error in firewall.go (5aecfa1c)
* - Fix some error in firewall.go (070605de)
* Add firewall option to the civogo (572557b6)
* fix test (ebde1c12)
* handle http errors (9c0f0967)
* add tests (37733dc1)
* Add functions to manipulate records (5f324bc9)
* - Fix a error network, now you can rename a network (79674176)
* Merge pull request #1 from civo/master (cbeb0cd4)
* change config for `form` (1ea969af)
* Add a default owner for the repo (fb92d21f)
* use Create instead of New (3055381e)
* New and delete Domain (5c6fe06b)
* add a result const (9c28c61c)
* fix GetDomain test (9df135c1)
* get and list records (5bf08246)
* update and new domain (acb6cd6a)
* - Update name in loadbalancer_test.go (4048f9c1)
* - Update the loadbalancer_test.go (fea3cf5b)
* - Fix some error in types in loadbalancer.go and loadbalancer_test.go (84b247e5)
* - Now can add, delete update and list the load balance (e2e8858b)
* - Initial files (4002274e)
* - Fix all recommendation from Andy, change NumTargetNodes from string to int, and in UpdateKubernetesClusters change map[string]string in favor of map[string]interface{} (2fdaa159)
* - Fix all recommendation from Andy (fa672090)
* - Final change in the kubernetes interface (7e0221f0)
* - Some change in kubernetes (9bb817f9)
* - First commit for kubernetes (05711724)
* - Update README.md with a list of all modules (58429347)
* Merge pull request #3 from civo/master (5a3cfcd6)
* Merge pull request #2 from civo/master (311f778e)
* - Update README.md, Fix error in loadbalancer.go (74877fa8)
* - Done the snapshot module (d2ede4b0)
* - First part of the snapshot api (eaf506c7)
* Merge branch 'master' into feature/api_snapshots (2408aaa3)
* Merge pull request #4 from civo/master (6d43cd65)
* Merge pull request #5 from civo/master (6e2137e8)
* - Add webhook module (467bd9a5)
* Merge pull request #6 from civo/master (55ee0710)
* - Fix some bug in dns.go dns_test.go (2edd5469)
* - Add option to update dns record (8ff8a494)
* Merge pull request #7 from civo/master (08c1c198)
* fix(loadbalancer): Fix data in test (29eab040)
* fix(loadbalancer): Fix struct in loadbalancer (450e0188)
* fix(firewall): Fix error in firewall test (f8658cf3)
* fix(firewall): Fix struct in firewall (6b4ca011)
* - Fix some bug in firewall.go and firewall_test.go (7f13e1c8)
* Merge pull request #9 from civo/master (27235987)
* Merge pull request #8 from civo/master (d118d875)
* fix(firewall): Change the id passed to the function (aff66676)
* fix(firewall): Fix name in the firewall (c480c528)
* fix(firewall): Fix name in the firewall (874f9ef6)
* feat(firewall): Add new option to search rule (58a86032)
* Merge pull request #11 from civo/master (53dc6a1e)
* Merge pull request #10 from alejandrojnm/hotfix/urgent_fix (99662c54)

0.2.1
=============
2020-03-27

* add template endpoints (c73d51fa)

0.2.0
=============
2020-03-24

* Minor tweaks to SSH key struct (d20f49e0)
* update the ssh key file (f5eab5e2)
* Add new option to ssh (fdf845a6)
* Fail the build if any lint issues (95ae34b1)

0.1.9
=============
2020-03-20

* add rename endpoint for firewall (588340b6)
* Updated the lib to use json and not form for all request (09617776)
* Add body to client error responses (a8ebeb95)
* Merge pull request #25 from alejandrojnm/hostfix/add_function_firewall (5ccea39d)
* Merge pull request #24 from alejandrojnm/hotfix/urgent_fix (7b07ee24)
* Change GetDNSRecord to find by ID not name (4f53b021)
* Remove DNSDomainID from DNSRecordConfig (edfa5583)
* Merge pull request #23 from alejandrojnm/feature/update_dns_record (d8c0e967)
* Fix instance test (5ae22ab1)
* Fix signature of SetInstanceTags (c470a7d5)
* Add GetInstanceConsoleURL method (1032a3be)
* Check links in go documentation (efbe8dbc)
* Change link to docs (83fdce86)
* Update badges (7cbcb38c)
* Fix tests after some minor refactoring (f149ccdc)
* Change GB to Gigabytes (f0fa3a15)
* Fix lint issues (62a82aec)
* Change workflow to custom action (143ac314)
* Move lint to its own action (541c21a6)
* Use user-contributed lint action (4e265ae4)
* Install golint before running it (0aa650f4)
* Add lint to tests and Go 1.14 (0bbdcd0c)
* Add changelog (f2cd42c5)
* Add Find* methods to all areas (d9686526)
* Fix some linting issues with Kubernetes (8dafd66a)
* Add LICENSE (dd906190)
* Return an error if creating a client with an empty API key (ee8dab3b)
* Remove progress table from README as its now module complete (1bfa0af0)
* Add minor tweaks to Webhooks (0471e7d2)
* Merge pull request #21 from alejandrojnm/feature/webhook (2d7d45db)
* Change minor naming issues in snapshot module (fc4a75b7)
* Merge pull request #20 from alejandrojnm/feature/api_snapshots (4b2ae45c)
* Fix a couple of comments (81e4652e)
* Add charge endpoint (457dfa5d)
* Change awkward names of units in field names (f20cab75)
* Add instance size endpoints (971a4a47)
* Add quota endpoint (7b54da2c)
* Add region endpoint (66b590cc)
* Change List functions to be plural and prefixing DNS (e3433bb0)
* Update display of README progress (43d6bc08)
* Merge pull request #19 from alejandrojnm/hotfix/loadbalance_readme (eacb35ca)
* Merge pull request #18 from alejandrojnm/master (41c7acb3)
* Merge pull request #16 from alejandrojnm/add-kubernetes-apis (730f2dd4)
* Add minor changes to Load Balancer methods (b8d5ddb8)
* Merge pull request #14 from alejandrojnm/feature/loadbalancer (e69c9746)
* Merge pull request #11 from rberrelleza/get-record (5590adda)
* Merge pull request #12 from rberrelleza/add-owners (338c4dfd)
* Add client.LastJSONResponse for the CLI to use (5378d41d)
* Merge pull request #7 from rberrelleza/fix-casing-forms (b3aba767)
* Merge pull request #9 from alejandrojnm/hostfix-network (f66cbcbb)
* Change string to map of strings (ab4ebaa8)
* Cleanup some minor issues (a69bef1d)
* Merge pull request #3 from rberrelleza/add-dns-apis (34870810)
* Merge pull request #4 from alejandrojnm/add-firewall-apis (554cc2b3)
* Merge pull request #5 from alejandrojnm/add-volumes-apis (9e8048fc)
* Merge pull request #6 from alejandrojnm/add-network-options (bbb04c3d)
* Change order of automated tests (57fae691)
* Ensure gofmt code compliance (e6d1cdcb)
* Remove panic (5d9bf4a7)
* Remove string interpolation, more go idiomatic this way (50fd2f10)
* Move JSON parsing for SimpleResponse to a helper method (8e8caf0a)
* Add CONTRIBUTING guide (a6d63538)
* Add status bage to README (1eabf12e)
* Add Go module support (2e043e15)
* Add GitHub action to test golang (e9c40745)
* Fix tests (08e7668c)
* Add Instance functions (264d4c87)
* Add Instance methods,first version (cd217b8b)
* Add tests for instance (31634d0f)
* Add Instance listing/fetching/creating (dc16cca8)
* - Add more options to the network options (5e258256)
* - Add New network options (6fbeb8a8)
* Add option to create a new network (e4a70af4)
* Fix error in volumes.go in the VolumesConfig (3db7ca46)
* Add volumes option to the civogo (52dca353)
* - Fix a error in firewall.go (5aecfa1c)
* - Fix some error in firewall.go (070605de)
* Add firewall option to the civogo (572557b6)
* fix test (ebde1c12)
* handle http errors (9c0f0967)
* add tests (37733dc1)
* Add functions to manipulate records (5f324bc9)
* - Fix a error network, now you can rename a network (79674176)
* Merge pull request #1 from civo/master (cbeb0cd4)
* change config for `form` (1ea969af)
* Add a default owner for the repo (fb92d21f)
* use Create instead of New (3055381e)
* New and delete Domain (5c6fe06b)
* add a result const (9c28c61c)
* fix GetDomain test (9df135c1)
* get and list records (5bf08246)
* update and new domain (acb6cd6a)
* - Update name in loadbalancer_test.go (4048f9c1)
* - Update the loadbalancer_test.go (fea3cf5b)
* - Fix some error in types in loadbalancer.go and loadbalancer_test.go (84b247e5)
* - Now can add, delete update and list the load balance (e2e8858b)
* - Initial files (4002274e)
* - Fix all recommendation from Andy, change NumTargetNodes from string to int, and in UpdateKubernetesClusters change map[string]string in favor of map[string]interface{} (2fdaa159)
* - Fix all recommendation from Andy (fa672090)
* - Final change in the kubernetes interface (7e0221f0)
* - Some change in kubernetes (9bb817f9)
* - First commit for kubernetes (05711724)
* - Update README.md with a list of all modules (58429347)
* Merge pull request #3 from civo/master (5a3cfcd6)
* Merge pull request #2 from civo/master (311f778e)
* - Update README.md, Fix error in loadbalancer.go (74877fa8)
* - Done the snapshot module (d2ede4b0)
* - First part of the snapshot api (eaf506c7)
* Merge branch 'master' into feature/api_snapshots (2408aaa3)
* Merge pull request #4 from civo/master (6d43cd65)
* Merge pull request #5 from civo/master (6e2137e8)
* - Add webhook module (467bd9a5)
* Merge pull request #6 from civo/master (55ee0710)
* - Fix some bug in dns.go dns_test.go (2edd5469)
* - Add option to update dns record (8ff8a494)
* Merge pull request #7 from civo/master (08c1c198)
* fix(loadbalancer): Fix data in test (29eab040)
* fix(loadbalancer): Fix struct in loadbalancer (450e0188)
* fix(firewall): Fix error in firewall test (f8658cf3)
* fix(firewall): Fix struct in firewall (6b4ca011)
* - Fix some bug in firewall.go and firewall_test.go (7f13e1c8)
* Merge pull request #9 from civo/master (27235987)
* Merge pull request #8 from civo/master (d118d875)
* fix(firewall): Change the id passed to the function (aff66676)
* fix(firewall): Fix name in the firewall (c480c528)
* fix(firewall): Fix name in the firewall (874f9ef6)
* feat(firewall): Add new option to search rule (58a86032)
* Merge pull request #11 from civo/master (53dc6a1e)
* Merge pull request #10 from alejandrojnm/hotfix/urgent_fix (99662c54)

0.1.8
=============
2020-03-19

* Updated the lib to use json and not form for all request (09617776)
* Add body to client error responses (a8ebeb95)

0.1.7
=============
2020-03-16

* Merge pull request #25 from alejandrojnm/hostfix/add_function_firewall (5ccea39d)

0.1.6
=============
2020-03-16

* Merge pull request #24 from alejandrojnm/hotfix/urgent_fix (7b07ee24)

0.1.5
=============
2020-03-12

* Change GetDNSRecord to find by ID not name (4f53b021)

0.1.4
=============
2020-03-12

* Remove DNSDomainID from DNSRecordConfig (edfa5583)
* Merge pull request #23 from alejandrojnm/feature/update_dns_record (d8c0e967)

0.1.3
=============
2020-03-10

* Fix instance test (5ae22ab1)

0.1.2
=============
2020-03-10

* Fix signature of SetInstanceTags (c470a7d5)

0.1.1
=============
2020-03-10

* Add GetInstanceConsoleURL method (1032a3be)
* Check links in go documentation (efbe8dbc)
* Change link to docs (83fdce86)
* Update badges (7cbcb38c)
* Fix tests after some minor refactoring (f149ccdc)
* Change GB to Gigabytes (f0fa3a15)
* Fix lint issues (62a82aec)
* Change workflow to custom action (143ac314)
* Move lint to its own action (541c21a6)
* Use user-contributed lint action (4e265ae4)
* Install golint before running it (0aa650f4)
* Add lint to tests and Go 1.14 (0bbdcd0c)

0.1.0
=============
2020-03-03

* Add changelog (f2cd42c5)
* Add Find* methods to all areas (d9686526)
* Fix some linting issues with Kubernetes (8dafd66a)
* Add LICENSE (dd906190)
* Return an error if creating a client with an empty API key (ee8dab3b)
* Remove progress table from README as its now module complete (1bfa0af0)
* Add minor tweaks to Webhooks (0471e7d2)
* Merge pull request #21 from alejandrojnm/feature/webhook (2d7d45db)
* Change minor naming issues in snapshot module (fc4a75b7)
* Merge pull request #20 from alejandrojnm/feature/api_snapshots (4b2ae45c)
* Fix a couple of comments (81e4652e)
* Add charge endpoint (457dfa5d)
* Change awkward names of units in field names (f20cab75)
* Add instance size endpoints (971a4a47)
* Add quota endpoint (7b54da2c)
* Add region endpoint (66b590cc)
* Change List functions to be plural and prefixing DNS (e3433bb0)
* Update display of README progress (43d6bc08)
* Merge pull request #19 from alejandrojnm/hotfix/loadbalance_readme (eacb35ca)
* Merge pull request #18 from alejandrojnm/master (41c7acb3)
* Merge pull request #16 from alejandrojnm/add-kubernetes-apis (730f2dd4)
* Add minor changes to Load Balancer methods (b8d5ddb8)
* Merge pull request #14 from alejandrojnm/feature/loadbalancer (e69c9746)
* Merge pull request #11 from rberrelleza/get-record (5590adda)
* Merge pull request #12 from rberrelleza/add-owners (338c4dfd)
* Add client.LastJSONResponse for the CLI to use (5378d41d)
* Merge pull request #7 from rberrelleza/fix-casing-forms (b3aba767)
* Merge pull request #9 from alejandrojnm/hostfix-network (f66cbcbb)
* Change string to map of strings (ab4ebaa8)
* Cleanup some minor issues (a69bef1d)
* Merge pull request #3 from rberrelleza/add-dns-apis (34870810)
* Merge pull request #4 from alejandrojnm/add-firewall-apis (554cc2b3)
* Merge pull request #5 from alejandrojnm/add-volumes-apis (9e8048fc)
* Merge pull request #6 from alejandrojnm/add-network-options (bbb04c3d)
* Change order of automated tests (57fae691)
* Ensure gofmt code compliance (e6d1cdcb)
* Remove panic (5d9bf4a7)
* Remove string interpolation, more go idiomatic this way (50fd2f10)
* Move JSON parsing for SimpleResponse to a helper method (8e8caf0a)
* Add CONTRIBUTING guide (a6d63538)
* Add status bage to README (1eabf12e)
* Add Go module support (2e043e15)
* Add GitHub action to test golang (e9c40745)
* Fix tests (08e7668c)
* Add Instance functions (264d4c87)
* Add Instance methods,first version (cd217b8b)
* Add tests for instance (31634d0f)
* Add Instance listing/fetching/creating (dc16cca8)
* - Add more options to the network options (5e258256)
* - Add New network options (6fbeb8a8)
* Add option to create a new network (e4a70af4)
* Fix error in volumes.go in the VolumesConfig (3db7ca46)
* Add volumes option to the civogo (52dca353)
* - Fix a error in firewall.go (5aecfa1c)
* - Fix some error in firewall.go (070605de)
* Add firewall option to the civogo (572557b6)
* fix test (ebde1c12)
* handle http errors (9c0f0967)
* add tests (37733dc1)
* Add functions to manipulate records (5f324bc9)
* - Fix a error network, now you can rename a network (79674176)
* Merge pull request #1 from civo/master (cbeb0cd4)
* change config for `form` (1ea969af)
* Add a default owner for the repo (fb92d21f)
* use Create instead of New (3055381e)
* New and delete Domain (5c6fe06b)
* add a result const (9c28c61c)
* fix GetDomain test (9df135c1)
* get and list records (5bf08246)
* update and new domain (acb6cd6a)
* - Update name in loadbalancer_test.go (4048f9c1)
* - Update the loadbalancer_test.go (fea3cf5b)
* - Fix some error in types in loadbalancer.go and loadbalancer_test.go (84b247e5)
* - Now can add, delete update and list the load balance (e2e8858b)
* - Initial files (4002274e)
* - Fix all recommendation from Andy, change NumTargetNodes from string to int, and in UpdateKubernetesClusters change map[string]string in favor of map[string]interface{} (2fdaa159)
* - Fix all recommendation from Andy (fa672090)
* - Final change in the kubernetes interface (7e0221f0)
* - Some change in kubernetes (9bb817f9)
* - First commit for kubernetes (05711724)
* - Update README.md with a list of all modules (58429347)
* Merge pull request #3 from civo/master (5a3cfcd6)
* Merge pull request #2 from civo/master (311f778e)
* - Update README.md, Fix error in loadbalancer.go (74877fa8)
* - Done the snapshot module (d2ede4b0)
* - First part of the snapshot api (eaf506c7)
* Merge branch 'master' into feature/api_snapshots (2408aaa3)
* Merge pull request #4 from civo/master (6d43cd65)
* Merge pull request #5 from civo/master (6e2137e8)
* - Add webhook module (467bd9a5)
* Merge pull request #6 from civo/master (55ee0710)


