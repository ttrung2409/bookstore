package com.bookstore.ecommerce.rest;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import com.bookstore.ecommerce.utils.NotFoundException;
import com.bookstore.ecommerce.utils.ThrowableFunc;
import org.springframework.http.ResponseEntity;
import lombok.var;

public abstract class ControllerBase {
  protected <T> ResponseEntity<?> executeQuery(ThrowableFunc<CompletableFuture<?>> query) {
    try {
      final var result = query.apply().get();

      return ResponseEntity.ok(result);
    } catch (ExecutionException e) {
      if (e.getCause() instanceof NotFoundException) {
        throw (NotFoundException) e.getCause();
      }

      return ResponseEntity.internalServerError().body(e);
    } catch (NotFoundException e) {
      return ResponseEntity.notFound().build();
    } catch (Exception e) {
      return ResponseEntity.internalServerError().body(e);
    }
  }

  protected <T> ResponseEntity<?> executeCommand(ThrowableFunc<CompletableFuture<?>> command) {
    try {
      final var result = command.apply().get();

      return ResponseEntity.ok(result);
    } catch (Exception e) {
      return ResponseEntity.internalServerError().body(e);
    }
  }
}

