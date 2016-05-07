# swarm-hooks

Docker Swarm provides a Docker Engine abstraction on top of a cluster of machines. This project, "Swarm-Hooks" exposes the Swarm API and wraps delegation with a pre/post hook processing chain. This pluggable chain allows the user to build clusters integrated with other "cluster-kernel" level services like secret managers, certificate authorities, monitoring subsystems, and fault injection tooling.

## Contributing

    # Get the code
    git clone https://github.com/buildertools/swarm-hooks.git

    # Start iterating
    make iterate

    # Do your work:
    #  Use the running service:
    #    open http://localhost:7890
    
    #  Watch the build logs:
    docker-compose logs service-iterate
    
    #  Watch the functional / integration tests:
    docker-compose logs integration
    
    #  Watch the service metrics:
    #    open http://localhost:3000
    
    # Stop working for the day:
    make stop
    
    # Produce a versioned release image
    make release

