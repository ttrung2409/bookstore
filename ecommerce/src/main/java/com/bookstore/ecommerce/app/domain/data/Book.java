package com.bookstore.ecommerce.app.domain.data;

import java.time.Instant;
import lombok.Builder;
import lombok.Getter;

@Builder
public class Book {
  @Getter
  private String id;
  @Getter
  private String title;
  @Getter
  private String subtitle;
  @Getter
  private String description;
  @Getter
  private String[] authors;
  @Getter
  private String publisher;
  @Getter
  private Instant publishedDate;
  @Getter
  private double averageRating;
  @Getter
  private int ratingsCount;
  @Getter
  private String thumbnailUrl;
  @Getter
  private String previewUrl;
  @Getter
  private int onhandQty;
  @Getter
  private int reservedQty;
}
