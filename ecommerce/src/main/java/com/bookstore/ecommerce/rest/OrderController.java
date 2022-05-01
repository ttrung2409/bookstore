package com.bookstore.ecommerce.rest;

import com.bookstore.ecommerce.app.order.command.Command;
import com.bookstore.ecommerce.app.order.command.CreateOrderRequest;
import com.bookstore.ecommerce.app.order.query.Query;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class OrderController extends ControllerBase {
  private final Query query;
  private final Command command;

  public OrderController(Query query, Command command) {
    this.query = query;
    this.command = command;
  }

  @GetMapping("/order/{id}")
  public ResponseEntity<?> get(@PathVariable String id) {
    return this.executeQuery(() -> {
      return this.query.getOrderDetails(id);
    });
  }

  @PostMapping("/order")
  public ResponseEntity<?> create(@RequestBody CreateOrderRequest request) {
    return this.executeCommand(() -> {
      return this.command.createOrder(request);
    });
  }
}
