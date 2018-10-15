class GoserverController < ApplicationController
  def index
    resp = Goserver.client[:goserver].call('GetImages')
    #logger.info resp.inspect

    @images = resp['Images']
    render json: @images
  end
end
