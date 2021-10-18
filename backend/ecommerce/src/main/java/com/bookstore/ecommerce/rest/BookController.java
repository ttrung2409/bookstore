package com.bookstore.ecommerce.rest;

import com.bookstore.ecommerce.app.order.query.Query;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class BookController extends ControllerBase {
  private final Query query;

  public BookController(Query query) {
    this.query = query;
  }

  @GetMapping("/book")
  public ResponseEntity<?> find(String term) {
    return this.executeQuery(() -> {
      return this.query.findBooks(term);
    });
  }
}
