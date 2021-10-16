package com.bookstore.ecommerce.app.order.query;

import java.util.Arrays;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Stream;

import com.bookstore.ecommerce.app.repository.BookRepository;
import com.ea.async.Async;

import org.springframework.stereotype.Component;

import lombok.experimental.ExtensionMethod;

@Component
@ExtensionMethod({ BookMapper.class })
public class BookQueryImpl implements BookQuery {
  private final BookRepository bookRepository;

  public BookQueryImpl(BookRepository bookRepository) {
    this.bookRepository = bookRepository;
  }

  @Override
  public CompletableFuture<Book[]> find(String term) {
    var books = Async.await(this.bookRepository.find(term));

    return Arrays.stream(books).map(book -> book.fromDataObject()).toArray();
  }
}
