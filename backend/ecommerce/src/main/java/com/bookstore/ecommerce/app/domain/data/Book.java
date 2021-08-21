package com.bookstore.ecommerce.app.domain.data;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;

@AllArgsConstructor
@Builder
public class Book {
  @Getter
  private String id;
  @Getter
  private String title;
  @Getter
  private String subTitle;
  @Getter
  private String description;
}
