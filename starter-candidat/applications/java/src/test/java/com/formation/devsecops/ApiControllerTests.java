package com.formation.devsecops;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class ApiControllerTests {

    @Autowired
    private ApiController controller;

    @Test
    void healthReturnsUp() {
        assertThat(controller.health()).containsEntry("status", "UP");
    }

    @Test
    void versionComesFromPom() {
        assertThat(controller.version()).containsEntry("version", "1.0.0");
    }

    @Test
    void infoContainsRuntimeDetails() {
        assertThat(controller.info()).containsKeys("application", "runtime", "hostname", "user");
    }
}
