package com.example.microservice.controller;

import com.example.microservice.model.UserData;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {

    @GetMapping("/user_a")
    public UserData getUserA() {
        return new UserData("User A", 30, "user_a@example.com");
    }

    @GetMapping("/user_b")
    public UserData getUserB() {
        return new UserData("User B", 25, "user_b@example.com");
    }
}
