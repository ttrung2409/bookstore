package com.bookstore.ecommerce.rest;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionException;
import com.bookstore.ecommerce.utils.NotFoundException;
import com.bookstore.ecommerce.utils.ThrowableSupplier;
import org.springframework.http.ResponseEntity;
import lombok.var;

public abstract class ControllerBase {
  protected ResponseEntity<?> executeQuery(ThrowableSupplier<CompletableFuture<?>> query) {
    try {
      final var result = query.get().join();

      return ResponseEntity.ok(result);
    } catch (CompletionException e) {
      throw e.getCause() instanceof NotFoundException
        ? (NotFoundException) e.getCause()
        : e.getCause() != null
          ? (RuntimeException) e
          : e;
    } catch (NotFoundException e) {
      return ResponseEntity.notFound().build();
    } catch (Exception e) {
      return ResponseEntity.internalServerError().body(e);
    }
  }

  protected ResponseEntity<?> executeCommand(ThrowableSupplier<CompletableFuture<?>> command) {
    try {
      final var result = command.get().join();

      return ResponseEntity.ok(result);
    } catch (CompletionException e) {
      throw e.getCause() != null ? (RuntimeException) e.getCause() : e;
    } catch (Exception e) {
      return ResponseEntity.internalServerError().body(e);
    }
  }
}

