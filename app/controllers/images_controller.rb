class ImagesController < ApplicationController
  # notice we've defined a method called index for a ImagesController instance. We tie this together with routes
  def index
    @images = Image.all # instance variables are prefixed with an @, to access images in the template

    respond_to do |format|
      format.html
      format.json { render json: @images }
    end
  end
end
