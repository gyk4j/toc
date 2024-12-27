# True Original Copy (TOC)

A backup and restore disaster recovery software to create and restore backups
between a local and remote site.

![TOC UI](../../wiki/assets/images/05-toc-progress.png)

> [!NOTE]
> Today, *the work is not yet finished* and TOC is still non-functional; TOC 
> remains a mock-up UI prototype with unfinished backend API services.
> 

*TOC* began life as a prototype as a testbed for various modern software 
technologies with the intention of learning and experimentation. It uses the 
following technologies:

- [Terraform][terraform] / [OpenTofu][opentofu]
- [Ansible][ansible]
- [Vagrant][vagrant]
- [Docker][docker], [Docker Compose][docker-compose], [s6 process supervisor][s6-overlay]
- [Go][golang], [Swagger][swagger]
- [React framework][react]
- [Material Design for Bootstrap (MDB)][mdb]

> [!CAUTION]
> Do **NOT** expect to use this for any real data &mdash; it **DOES NOT** work.
>

To learn more, please visit the [project wiki][wiki].

[terraform]: https://www.terraform.io/
[opentofu]: https://opentofu.org/
[ansible]: https://www.ansible.com/
[vagrant]: https://www.vagrantup.com/
[docker]: https://www.docker.com/community/open-source/
[docker-compose]: https://docs.docker.com/compose/
[s6-overlay]: https://skarnet.org/software/s6/
[golang]: https://go.dev/
[swagger]: https://swagger.io/
[react]: https://react.dev/
[mdb]: https://mdbootstrap.com/
[wiki]: ../../wiki/Home
